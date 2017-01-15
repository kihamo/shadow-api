package api

import (
	"github.com/kihamo/shadow/components/metrics"
)

const (
	MetricApiProcedureExecuteTime = "api.procedure.execute_time"
)

var (
	metricApiProcedureExecuteTime metrics.Timer
)

func (c *Component) MetricsRegister(m *metrics.Component) {
	metricApiProcedureExecuteTime = m.NewTimer(MetricApiProcedureExecuteTime)
}
