package internal

import (
	"github.com/kihamo/shadow-api/components/api"
	"github.com/kihamo/shadow-api/components/api/internal/procedures"
)

func (c *Component) GetApiProcedures() []api.Procedure {
	return []api.Procedure{
		&procedures.VersionProcedure{
			Application: c.application,
		},
	}
}
