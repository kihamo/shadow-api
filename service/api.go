package service

func (s *ApiService) GetApiProcedures() []ApiProcedure {
	return []ApiProcedure{
		&PingProcedure{},
		&VersionProcedure{},
	}
}
