package api

import (
	"fmt"
	"strings"

	"github.com/kihamo/shadow/components/dashboard"
)

type IndexHandler struct {
	dashboard.Handler

	component *Component
}

func (h *IndexHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	host := r.Config().GetString(ConfigHost)
	if host == "0.0.0.0" && r.Original().Host != "" {
		s := strings.Split(r.Original().Host, ":")
		host = s[0]
	}

	h.Render(r.Context(), ComponentName, "index", map[string]interface{}{
		"apiUrl":     fmt.Sprintf("ws://%s:%d/", host, r.Config().GetInt64(ConfigPort)),
		"procedures": h.component.GetProcedures(),
	})
}
