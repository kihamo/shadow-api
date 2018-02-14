package internal

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-api/components/api"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/logger"
	"github.com/kihamo/shadow/components/metrics"
	"gopkg.in/jcelliott/turnpike.v2"
)

type Component struct {
	application shadow.Application
	config      config.Component
	logger      logger.Logger
	routes      []dashboard.Route

	turnpikeLogger *Logger
	procedures     []api.Procedure
	metricEnabled  bool
}

func (c *Component) Name() string {
	return api.ComponentName
}

func (c *Component) Version() string {
	return api.ComponentVersion
}

func (c *Component) Dependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: logger.ComponentName,
		},
		{
			Name: metrics.ComponentName,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.application = a
	c.config = a.GetComponent(config.ComponentName).(config.Component)
	c.metricEnabled = a.HasComponent(metrics.ComponentName)

	return nil
}

func (c *Component) Run(wg *sync.WaitGroup) error {
	c.logger = logger.NewOrNop(c.Name(), c.application)

	c.turnpikeLogger = NewLogger(c.logger)
	if c.config.Bool(api.ConfigLoggingEnabled) {
		c.turnpikeLogger.On()
	} else {
		c.turnpikeLogger.Off()
	}

	turnpike.SetLogger(c.turnpikeLogger)

	handler := turnpike.NewBasicWebsocketServer(c.Name())
	client, err := handler.GetLocalClient(c.Name(), nil)
	if err != nil {
		return err
	}

	components, err := c.application.GetComponents()
	if err != nil {
		return err
	}

	for _, cmp := range components {
		if cmpApi, ok := cmp.(api.HasApiProcedures); ok {
			for _, procedure := range cmpApi.GetApiProcedures() {
				name := procedure.GetName()

				if c.HasProcedure(name) {
					c.logger.Warn("Procedure already exists. Ignore procedure.", map[string]interface{}{
						"procedure": name,
						"service":   cmp.Name(),
					})

					continue
				}

				procedureWrapper := func(procedure api.Procedure) turnpike.BasicMethodHandler {
					return func(args []interface{}, kwargs map[string]interface{}) *turnpike.CallResult {
						beforeTime := time.Now()
						defer func() {
							if c.metricEnabled {
								metricExecuteTime.UpdateSince(beforeTime)
								metricExecuteTime.With("procedure", procedure.GetName()).UpdateSince(beforeTime)
							}
						}()

						if autoValidation, ok := procedure.(api.ProcedureWithRequest); ok {
							out := autoValidation.GetRequest()
							request := NewRequest(out, args, kwargs)
							if ok, errors := request.Valid(); !ok {
								return &turnpike.CallResult{
									Kwargs: map[string]interface{}{
										"errors": errors,
									},
									Err: turnpike.URI(api.ErrorInvalidArgument),
								}
							}

							return autoValidation.Run(out)
						}

						if simple, ok := procedure.(api.ProcedureSimple); ok {
							return simple.Run(args, kwargs)
						}

						c.logger.Error("Error procedure interace", map[string]interface{}{
							"procedure": name,
							"service":   cmp.Name(),
							"error":     err.Error(),
						})

						return &turnpike.CallResult{
							Err: turnpike.URI(api.ErrorUnknownProcedure),
						}
					}
				}

				if err = client.BasicRegister(name, procedureWrapper(procedure)); err != nil {
					c.logger.Error("Error register api procedure", map[string]interface{}{
						"procedure": name,
						"service":   cmp.Name(),
						"error":     err.Error(),
					})
					// ignore error
				} else {
					c.logger.Debug("Register procedure", map[string]interface{}{
						"procedure": name,
						"service":   cmp.Name(),
					})
				}
				c.procedures = append(c.procedures, procedure)
			}
		}
	}

	go func(handler *turnpike.WebsocketServer) {
		defer wg.Done()

		// TODO: ssl

		addr := fmt.Sprintf("%s:%d", c.config.String(api.ConfigHost), c.config.Int64(api.ConfigPort))

		c.logger.Info("Running service", map[string]interface{}{
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

			if c.config.Bool(api.ConfigLoggingEnabled) {
				c.logger.Infof("Connection from %s", r.RemoteAddr)
			}

			handler.ServeHTTP(w, r)
		})

		if err := server.ListenAndServe(); err != nil {
			c.logger.Fatalf("Could not start api [%d]: %s\n", os.Getpid(), err.Error())
		}
	}(handler)

	return nil
}

func (c *Component) GetProcedures() []api.Procedure {
	return c.procedures
}

func (c *Component) GetProcedure(procedure string) api.Procedure {
	for _, p := range c.procedures {
		if p.GetName() == procedure {
			return p
		}
	}

	return nil
}

func (c *Component) HasProcedure(procedure string) bool {
	return c.GetProcedure(procedure) != nil
}

func (c *Component) GetClient() (*turnpike.Client, error) {
	addr := fmt.Sprintf("ws://%s:%d/", c.config.String("api.host"), c.config.Int64("api.port"))

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
