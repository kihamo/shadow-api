package handlers

import (
	"fmt"
	"strings"

	"github.com/kihamo/shadow-api/components/api"
	"github.com/kihamo/shadow/components/dashboard"
)

type ManagerHandler struct {
	dashboard.Handler

	Component api.Component
}

func (h *ManagerHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	host := r.Config().String(api.ConfigHost)
	if host == "0.0.0.0" && r.Original().Host != "" {
		s := strings.Split(r.Original().Host, ":")
		host = s[0]
	}

	h.Render(r.Context(), h.Component.Name(), "manager", map[string]interface{}{
		"apiUrl":     fmt.Sprintf("ws://%s:%d/", host, r.Config().Int64(api.ConfigPort)),
		"procedures": h.Component.GetProcedures(),
	})
}
