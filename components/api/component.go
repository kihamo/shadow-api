package api

import (
	"github.com/kihamo/shadow"
	"gopkg.in/jcelliott/turnpike.v2"
)

type Component interface {
	shadow.Component

	GetProcedures() []Procedure
	GetProcedure(procedure string) Procedure
	HasProcedure(procedure string) bool
	GetClient() (*turnpike.Client, error)
}
