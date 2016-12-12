package service

import (
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow/service/frontend"
)

func (s *ApiService) GetTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "templates",
	}
}

func (s *ApiService) GetFrontendMenu() *frontend.FrontendMenu {
	return &frontend.FrontendMenu{
		Name: "Api",
		Url:  "/api",
		Icon: "exchange",
	}
}

func (s *ApiService) SetFrontendHandlers(router *frontend.Router) {
	router.ServeFiles("/js/api/*filepath", &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "public/js",
	})

	router.GET(s, "/api", &IndexHandler{
		service: s,
		config:  s.config,
	})
}
