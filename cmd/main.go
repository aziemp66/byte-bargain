package main

import (
	"fmt"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"

	envCommon "github.com/aziemp66/byte-bargain/common/env"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"
)

func main() {
	cfg := envCommon.LoadConfig()
	goviewConfig := goview.Config{
		Root:         "webviews",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: false,
		Delims: goview.Delims{
			Left:  "{{",
			Right: "}}",
		},
	}

	httpServer := httpCommon.NewHTTPServer(cfg.GinMode)

	httpServer.Router.HTMLRender = ginview.New(goviewConfig)

	sessionManager := sessionCommon.NewSessionManager([]byte(cfg.AccessTokenKey))

	httpServer.Router.Use(sessionManager.GetSessionHandler())
	httpServer.Router.Use(httpCommon.MiddlewareErrorHandler())

	err := httpServer.Router.Run(fmt.Sprintf(":%d", cfg.Port))

	if err != nil {
		panic(err)
	}
}
