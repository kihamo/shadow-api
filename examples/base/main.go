package main // import "github.com/kihamo/shadow-api/examples/base"

import (
	"log"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-api/components/api"
	"github.com/kihamo/shadow/components/alerts"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/logger"
	"github.com/kihamo/shadow/components/metrics"
)

func main() {
	application, err := shadow.NewApp(
		[]shadow.Component{
			new(config.Component),
			new(logger.Component),
			new(metrics.Component),
			new(alerts.Component),
			new(dashboard.Component),
			new(api.Component),
		},
		"Api",
		"1.0",
		"12345-full",
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = application.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
