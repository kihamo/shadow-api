package api

import (
	"github.com/kihamo/shadow/components/config"
)

const (
	ConfigApiHost          = "api.host"
	ConfigApiPort          = "api.port"
	ConfigApiSecureEnabled = "api.secure.enabled"
	ConfigApiSecureCrt     = "api.secure.crt"
	ConfigApiSecureKey     = "api.secure.key"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:     ConfigApiHost,
			Default: "localhost",
			Usage:   "API socket host",
			Type:    config.ValueTypeString,
		},
		{
			Key:     ConfigApiPort,
			Default: 8001,
			Usage:   "API socket port",
			Type:    config.ValueTypeInt,
		},
		{
			Key:     ConfigApiSecureEnabled,
			Default: false,
			Usage:   "API enable SSL",
			Type:    config.ValueTypeBool,
		},
		{
			Key:   ConfigApiSecureCrt,
			Usage: "API path to SSL crt file",
			Type:  config.ValueTypeString,
		},
		{
			Key:   ConfigApiSecureKey,
			Usage: "API path to SSL key file",
			Type:  config.ValueTypeString,
		},
	}
}
