package internal

import (
	"github.com/kihamo/shadow-api/components/api"
	"github.com/kihamo/snitch"
)

const (
	MetricExecuteTime = api.ComponentName + "_request_duration_seconds"
)

var (
	metricExecuteTime = snitch.NewTimer(MetricExecuteTime, "Total request duration")
)

func (c *Component) Metrics() snitch.Collector {
	return metricExecuteTime
}
