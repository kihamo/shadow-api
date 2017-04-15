package api

import (
	"github.com/kihamo/snitch"
)

const (
	MetricExecuteTime          = ComponentName + ".execute_time"
	MetricProcedureExecuteTime = ComponentName + ".procedure.execute_time"
)

var (
	metricExecuteTime snitch.Timer
)

func (c *Component) Metrics() snitch.Collector {
	metricExecuteTime = snitch.NewTimer(MetricExecuteTime)

	return metricExecuteTime
}
