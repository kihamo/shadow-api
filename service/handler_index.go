package service

import (
	"fmt"
	"strings"

	"github.com/kihamo/shadow/service/frontend"
)

type IndexHandler struct {
	frontend.AbstractFrontendHandler
}

func (h *IndexHandler) Handle() {
	h.SetTemplate("index.tpl.html")
	h.SetPageTitle("Api")
	h.SetPageHeader("Api")

	service := h.Service.(*ApiService)

	host := service.config.GetString("api.host")
	if host == "0.0.0.0" && h.Input.Host != "" {
		s := strings.Split(h.Input.Host, ":")
		host = s[0]
	}

	h.SetVar("ApiUrl", fmt.Sprintf("ws://%s:%d/", host, service.config.GetInt64("api.port")))
	h.SetVar("Procedures", service.GetProcedures())
}
