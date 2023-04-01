package main

import (
	"fmt"

	envCommon "github.com/aziemp66/byte-bargain/common/env"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"
)

func main() {
	cfg := envCommon.LoadConfig()

	httpServer := httpCommon.NewHTTPServer(cfg.GinMode)

	sessionManager := sessionCommon.NewSessionManager([]byte(cfg.AccessTokenKey))

	httpServer.Router.Use(sessionManager.GetSessionHandler())
	httpServer.Router.Use(httpCommon.MiddlewareErrorHandler())

	err := httpServer.Router.Run(fmt.Sprintf(":%d", cfg.Port))

	if err != nil {
		panic(err)
	}
}
