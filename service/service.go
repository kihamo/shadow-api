package service

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/resource/logger"
	"github.com/kihamo/shadow/resource/metrics"
	"gopkg.in/jcelliott/turnpike.v2"
)

type ServiceApiHandler interface {
	GetApiProcedures() []ApiProcedure
}

type ApiService struct {
	application *shadow.Application
	config      *config.Resource
	logger      logger.Logger

	procedures []ApiProcedure
}

func (s *ApiService) GetName() string {
	return "api"
}

func (s *ApiService) Init(a *shadow.Application) error {
	s.application = a

	resourceConfig, err := a.GetResource("config")
	if err != nil {
		return err
	}
	s.config = resourceConfig.(*config.Resource)

	return nil
}

func (s *ApiService) Run(wg *sync.WaitGroup) error {
	if resourceLogger, err := s.application.GetResource("logger"); err == nil {
		s.logger = resourceLogger.(*logger.Resource).Get(s.GetName())
	} else {
		s.logger = logger.NopLogger
	}

	turnpike.SetLogger(s.logger)

	handler := turnpike.NewBasicWebsocketServer(s.GetName())
	client, err := handler.GetLocalClient(s.GetName(), nil)
	if err != nil {
		return err
	}

	for _, service := range s.application.GetServices() {
		if serviceCast, ok := service.(ServiceApiHandler); ok {
			for _, procedure := range serviceCast.GetApiProcedures() {
				name := procedure.GetName()

				if s.HasProcedure(name) {
					s.logger.Warn("Procedure already exists. Ignore procedure.", map[string]interface{}{
						"procedure": name,
						"service":   service.GetName(),
					})

					continue
				}

				if err := procedure.Init(service, s.application); err != nil {
					s.logger.Error("Procedure init failed", map[string]interface{}{
						"procedure": name,
						"service":   service.GetName(),
						"error":     err.Error(),
					})

					return err
				}

				procedureWrapper := func(procedure ApiProcedure) turnpike.BasicMethodHandler {
					var procedureMetricApiProcedureExecuteTime metrics.Timer
					if metricApiProcedureExecuteTime != nil {
						procedureMetricApiProcedureExecuteTime = metricApiProcedureExecuteTime.With("procedure", procedure.GetName())
					}

					return func(args []interface{}, kwargs map[string]interface{}) *turnpike.CallResult {
						beforeTime := time.Now()
						defer func() {
							if procedureMetricApiProcedureExecuteTime != nil {
								procedureMetricApiProcedureExecuteTime.ObserveDurationByTime(beforeTime)
							}
						}()

						if autoValidation, ok := procedure.(ApiProcedureRequest); ok {
							out := autoValidation.GetRequest()
							request := NewRequest(out, args, kwargs)
							if ok, errors := request.Valid(); !ok {
								return &turnpike.CallResult{
									Kwargs: map[string]interface{}{
										"errors": errors,
									},
									Err: turnpike.URI(ErrorInvalidArgument),
								}
							}

							return autoValidation.Run(out)
						}

						if simple, ok := procedure.(ApiProcedureSimple); ok {
							return simple.Run(args, kwargs)
						}

						s.logger.Error("Error procedure interace", map[string]interface{}{
							"procedure": name,
							"service":   service.GetName(),
							"error":     err.Error(),
						})

						return &turnpike.CallResult{
							Err: turnpike.URI(ErrorUnknownProcedure),
						}
					}
				}

				if err = client.BasicRegister(name, procedureWrapper(procedure)); err != nil {
					s.logger.Error("Error register api procedure", map[string]interface{}{
						"procedure": name,
						"service":   service.GetName(),
						"error":     err.Error(),
					})
					// ignore error
				} else {
					s.logger.Debug("Register procedure", map[string]interface{}{
						"procedure": name,
						"service":   service.GetName(),
					})
				}
				s.procedures = append(s.procedures, procedure)
			}
		}
	}

	go func(handler *turnpike.WebsocketServer) {
		defer wg.Done()

		// TODO: ssl

		addr := fmt.Sprintf("%s:%d", s.config.GetString("api.host"), s.config.GetInt64("api.port"))

		s.logger.Info("Running service", map[string]interface{}{
			"addr": addr,
			"pid":  os.Getpid(),
		})

		mux := http.NewServeMux()
		server := &http.Server{
			Handler: mux,
			Addr:    addr,
		}

		mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// FiXME: Magic
			delete(r.Header, "Origin")

			s.logger.Infof("Connection from %s", r.RemoteAddr)
			handler.ServeHTTP(w, r)
		})

		if err := server.ListenAndServe(); err != nil {
			s.logger.Fatalf("Could not start api [%d]: %s\n", os.Getpid(), err.Error())
		}
	}(handler)

	return nil
}

func (s *ApiService) GetProcedures() []ApiProcedure {
	return s.procedures
}

func (s *ApiService) GetProcedure(procedure string) ApiProcedure {
	for _, p := range s.procedures {
		if p.GetName() == procedure {
			return p
		}
	}

	return nil
}

func (s *ApiService) HasProcedure(procedure string) bool {
	return s.GetProcedure(procedure) != nil
}

func (s *ApiService) GetClient() (*turnpike.Client, error) {
	addr := fmt.Sprintf("ws://%s:%d/", s.config.GetString("api.host"), s.config.GetInt64("api.port"))

	client, err := turnpike.NewWebsocketClient(turnpike.JSON, addr, nil, nil)
	if err != nil {
		return nil, err
	}

	_, err = client.JoinRealm("api", nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
