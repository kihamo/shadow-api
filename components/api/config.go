package api

import (
	"github.com/kihamo/shadow/components/config"
)

const (
	ConfigApiHost           = "api.host"
	ConfigApiPort           = "api.port"
	ConfigApiLoggingEnabled = "api.logging.enable"
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
			Key:      ConfigApiLoggingEnabled,
			Default:  true,
			Usage:    "API enable logging",
			Type:     config.ValueTypeBool,
			Editable: true,
		},
	}
}

func (c *Component) GetConfigWatchers() map[string][]config.Watcher {
	return map[string][]config.Watcher{
		ConfigApiLoggingEnabled: {c.watchApiLoggingEnabled},
	}
}

func (c *Component) watchApiLoggingEnabled(_ string, newValue interface{}, _ interface{}) {
	if newValue.(bool) {
		c.turnpikeLogger.On()
	} else {
		c.turnpikeLogger.Off()
	}
}
