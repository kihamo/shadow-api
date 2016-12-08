package service

import (
	"fmt"
	"os"

	"github.com/kihamo/shadow/resource/config"
)

func (s *ApiService) GetConfigVariables() []config.Variable {
	pathCrt := ""
	pathKey := ""

	dir, err := os.Getwd()
	if err == nil {
		pathCrt = fmt.Sprint(dir, "/server.crt")
		pathKey = fmt.Sprint(dir, "/server.key")
	}

	return []config.Variable{
		{
			Key:   "api.host",
			Value: "0.0.0.0",
			Usage: "API socket host",
		},
		{
			Key:   "api.port",
			Value: 8001,
			Usage: "API socket port",
		},
		{
			Key:   "api.secure",
			Value: false,
			Usage: "API enable SSL",
		},
		{
			Key:   "api.secure-crt",
			Value: pathCrt,
			Usage: "API path to SSL crt file",
		},
		{
			Key:   "api.secure-key",
			Value: pathKey,
			Usage: "API path to SSL key file",
		},
	}
}
