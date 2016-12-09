package service

import (
	"errors"

	"github.com/kihamo/shadow"
	"gopkg.in/jcelliott/turnpike.v2"
)

const (
	ErrorUnknownProcedure = "api.unknown-procedure"
	ErrorInvalidArgument  = "api.invalid-argument"
)

type ApiProcedure interface {
	Init(shadow.Service, *shadow.Application) error
	GetName() string
}

type ApiProcedureSimple interface {
	ApiProcedure
	Run([]interface{}, map[string]interface{}) *turnpike.CallResult
}

type ApiProcedureRequest interface {
	ApiProcedure
	GetRequest() interface{}
	Run(interface{}) *turnpike.CallResult
}

type AbstractApiProcedure struct {
	ApiProcedure
	Application *shadow.Application
	Service     shadow.Service
	ApiService  *ApiService
}

func (p *AbstractApiProcedure) Init(s shadow.Service, a *shadow.Application) error {
	p.Application = a
	p.Service = s

	if apiService, err := a.GetService("api"); err == nil {
		if castService, ok := apiService.(*ApiService); ok {
			p.ApiService = castService
			return nil
		}
	}

	return errors.New("Api service not found")
}

func (p *AbstractApiProcedure) GetResult(args []interface{}, kwargs map[string]interface{}) *turnpike.CallResult {
	return &turnpike.CallResult{
		Args:   args,
		Kwargs: kwargs,
	}
}

func (p *AbstractApiProcedure) GetError(err string) *turnpike.CallResult {
	return &turnpike.CallResult{
		Err: turnpike.URI(err),
	}
}
