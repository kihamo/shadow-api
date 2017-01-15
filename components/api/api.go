package api

func (c *Component) GetApiProcedures() []ApiProcedure {
	return []ApiProcedure{
		&PingProcedure{},
		&VersionProcedure{
			version: c.application.GetVersion(),
			build:   c.application.GetBuild(),
		},
	}
}
