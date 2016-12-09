package main // import "github.com/kihamo/shadow-api/examples/base"

import (
	"log"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-api/service"
	"github.com/kihamo/shadow/resource/alerts"
	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/resource/logger"
	"github.com/kihamo/shadow/resource/metrics"
	"github.com/kihamo/shadow/resource/template"
	"github.com/kihamo/shadow/service/frontend"
	"github.com/kihamo/shadow/service/system"
)

var (
	Version = "0.0"
	Build   = "0-0000000"
)

func main() {
	application, err := shadow.NewApplication(
		[]shadow.Resource{
			new(config.Resource),
			new(logger.Resource),
			new(metrics.Resource),
			new(template.Resource),
			new(alerts.Resource),
		},
		[]shadow.Service{
			new(system.SystemService),
			new(frontend.FrontendService),
			new(service.ApiService),
		},
		"Api",
		Version,
		Build,
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = application.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
