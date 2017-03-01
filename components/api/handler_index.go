package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kihamo/shadow/components/dashboard"
)

type IndexHandler struct {
	dashboard.Handler

	component *Component
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	config := dashboard.ConfigFromContext(r.Context())

	host := config.GetString(ConfigApiHost)
	if host == "0.0.0.0" && r.Host != "" {
		s := strings.Split(r.Host, ":")
		host = s[0]
	}

	h.Render(r.Context(), ComponentName, "index", map[string]interface{}{
		"apiUrl":     fmt.Sprintf("ws://%s:%d/", host, config.GetInt64(ConfigApiPort)),
		"procedures": h.component.GetProcedures(),
	})
}
