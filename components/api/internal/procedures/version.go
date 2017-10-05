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
		"build_date": p.Application.GetBuildDate(),
		"build":      p.Application.GetBuild(),
		"start_date": p.Application.GetStartDate(),
		"uptime":     p.Application.GetUptime(),
		"version":    p.Application.GetVersion(),
	})
}
