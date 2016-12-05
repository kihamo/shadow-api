package main // import "github.com/kihamo/shadow-api/cmd/example"

import (
	"log"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-api/service"
	"github.com/kihamo/shadow/resource"
	"github.com/kihamo/shadow/resource/alerts"
	"github.com/kihamo/shadow/resource/metrics"
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
			new(resource.Config),
			new(resource.Logger),
			new(resource.Template),
			new(alerts.Alerts),
			new(metrics.Metrics),
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
