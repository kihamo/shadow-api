package service

import (
	"github.com/kihamo/shadow/resource/metrics"
)

const (
	MetricApiProcedureExecuteTime = "api.procedure.execute_time"
)

var (
	metricApiProcedureExecuteTime metrics.Timer
)

func (r *ApiService) MetricsRegister(m *metrics.Resource) {
	metricApiProcedureExecuteTime = m.NewTimer(MetricApiProcedureExecuteTime)
}
