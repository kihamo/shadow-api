package api

import (
	"gopkg.in/jcelliott/turnpike.v2"
)

type HasApiProcedures interface {
	GetApiProcedures() []Procedure
}

type Procedure interface {
	GetName() string
}

type ProcedureSimple interface {
	Procedure
	Run([]interface{}, map[string]interface{}) *turnpike.CallResult
}

type ProcedureWithRequest interface {
	Procedure
	GetRequest() interface{}
	Run(interface{}) *turnpike.CallResult
}

type ProcedureBase struct {
	Procedure
}

func (p *ProcedureBase) GetResult(args []interface{}, kwargs map[string]interface{}) *turnpike.CallResult {
	return &turnpike.CallResult{
		Args:   args,
		Kwargs: kwargs,
	}
}

func (p *ProcedureBase) GetError(err string) *turnpike.CallResult {
	return &turnpike.CallResult{
		Err: turnpike.URI(err),
	}
}
