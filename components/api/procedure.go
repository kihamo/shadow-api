package api

import (
	"gopkg.in/jcelliott/turnpike.v2"
)

const (
	ErrorUnknownProcedure = "api.unknown-procedure"
	ErrorInvalidArgument  = "api.invalid-argument"
)

type ApiProcedure interface {
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

type Procedure struct {
	ApiProcedure
}

func (p *Procedure) GetResult(args []interface{}, kwargs map[string]interface{}) *turnpike.CallResult {
	return &turnpike.CallResult{
		Args:   args,
		Kwargs: kwargs,
	}
}

func (p *Procedure) GetError(err string) *turnpike.CallResult {
	return &turnpike.CallResult{
		Err: turnpike.URI(err),
	}
}
