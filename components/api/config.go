package api

import (
	"github.com/kihamo/shadow/components/config"
)

const (
	ConfigHost           = ComponentName + ".host"
	ConfigPort           = ComponentName + ".port"
	ConfigLoggingEnabled = ComponentName + ".logging.enable"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:     ConfigHost,
			Default: "localhost",
			Usage:   "API socket host",
			Type:    config.ValueTypeString,
		},
		{
			Key:     ConfigPort,
			Default: 8001,
			Usage:   "API socket port",
			Type:    config.ValueTypeInt,
		},
		{
			Key:      ConfigLoggingEnabled,
			Default:  true,
			Usage:    "API enable logging",
			Type:     config.ValueTypeBool,
			Editable: true,
		},
	}
}

func (c *Component) GetConfigWatchers() map[string][]config.Watcher {
	return map[string][]config.Watcher{
		ConfigLoggingEnabled: {c.watchLoggingEnabled},
	}
}

func (c *Component) watchLoggingEnabled(_ string, newValue interface{}, _ interface{}) {
	if newValue.(bool) {
		c.turnpikeLogger.On()
	} else {
		c.turnpikeLogger.Off()
	}
}
