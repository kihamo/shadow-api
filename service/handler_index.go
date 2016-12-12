package service

import (
	"fmt"
	"strings"

	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/service/frontend"
)

type IndexHandler struct {
	frontend.AbstractFrontendHandler

	service *ApiService
	config  *config.Resource
}

func (h *IndexHandler) Handle() {
	h.SetTemplate("index.tpl.html")
	h.SetPageTitle("Api")
	h.SetPageHeader("Api")

	host := h.config.GetString(ConfigApiHost)
	if host == "0.0.0.0" && h.Input.Host != "" {
		s := strings.Split(h.Input.Host, ":")
		host = s[0]
	}

	h.SetVar("ApiUrl", fmt.Sprintf("ws://%s:%d/", host, h.config.GetInt64(ConfigApiPort)))
	h.SetVar("Procedures", h.service.GetProcedures())
}
