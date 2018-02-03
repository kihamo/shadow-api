package procedures

import (
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-api/components/api"
	"gopkg.in/jcelliott/turnpike.v2"
)

type VersionProcedure struct {
	api.ProcedureBase

	Application shadow.Application
}

func (p *VersionProcedure) GetName() string {
	return "api.version"
}

func (p *VersionProcedure) Run([]interface{}, map[string]interface{}) *turnpike.CallResult {
	return p.GetResult(nil, map[string]interface{}{
		"name":       p.Application.Name(),
		"version":    p.Application.Version(),
		"build":      p.Application.Build(),
		"build_date": p.Application.BuildDate(),
		"start_date": p.Application.StartDate(),
		"uptime":     p.Application.Uptime(),
	})
}
