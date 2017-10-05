package procedures

import (
	"github.com/kihamo/shadow-api/components/api"
	"gopkg.in/jcelliott/turnpike.v2"
)

type PingProcedure struct {
	api.ProcedureBase
}

func (p *PingProcedure) GetName() string {
	return "api.ping"
}

func (p *PingProcedure) Run([]interface{}, map[string]interface{}) *turnpike.CallResult {
	return p.GetResult([]interface{}{"pong"}, nil)
}
