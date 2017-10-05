package internal

import (
	"github.com/kihamo/shadow-api/components/api"
	"github.com/kihamo/shadow/components/config"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(
			api.ConfigHost,
			config.ValueTypeString,
			"localhost",
			"API socket host",
			false,
			nil,
			nil),
		config.NewVariable(
			api.ConfigPort,
			config.ValueTypeInt,
			8001,
			"API socket port",
			false,
			nil,
			nil),
		config.NewVariable(
			api.ConfigLoggingEnabled,
			config.ValueTypeBool,
			true,
			"API enable logging",
			true,
			nil,
			nil),
	}
}

func (c *Component) GetConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher(c.GetName(), []string{api.ConfigLoggingEnabled}, c.watchLoggingEnabled),
	}
}

func (c *Component) watchLoggingEnabled(_ string, newValue interface{}, _ interface{}) {
	if newValue.(bool) {
		c.turnpikeLogger.On()
	} else {
		c.turnpikeLogger.Off()
	}
}
