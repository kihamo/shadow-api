package internal

import (
	"github.com/kihamo/shadow-api/components/api"
	"github.com/kihamo/shadow/components/config"
)

func (c *Component) ConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(
			api.ConfigHost,
			config.ValueTypeString,
			"localhost",
			"API socket host",
			false,
			"Listen",
			nil,
			nil),
		config.NewVariable(
			api.ConfigPort,
			config.ValueTypeInt,
			8001,
			"API socket port",
			false,
			"Listen",
			nil,
			nil),
		config.NewVariable(
			api.ConfigLoggingEnabled,
			config.ValueTypeBool,
			true,
			"API enable logging",
			true,
			"Others",
			nil,
			nil),
	}
}

func (c *Component) ConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher(c.Name(), []string{api.ConfigLoggingEnabled}, c.watchLoggingEnabled),
	}
}

func (c *Component) watchLoggingEnabled(_ string, newValue interface{}, _ interface{}) {
	if newValue.(bool) {
		c.turnpikeLogger.On()
	} else {
		c.turnpikeLogger.Off()
	}
}
