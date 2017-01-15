package api

import (
	"gopkg.in/jcelliott/turnpike.v2"
)

type VersionProcedure struct {
	Procedure

	version string
	build   string
}

func (p *VersionProcedure) GetName() string {
	return "api.version"
}

func (p *VersionProcedure) Run([]interface{}, map[string]interface{}) *turnpike.CallResult {
	return p.GetResult(nil, map[string]interface{}{
		"version": p.version,
		"build":   p.build,
	})
}
