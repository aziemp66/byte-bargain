package main

import (
	"fmt"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"gopkg.in/gomail.v2"

	dbCommon "github.com/aziemp66/byte-bargain/common/db"
	envCommon "github.com/aziemp66/byte-bargain/common/env"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	jwtCommon "github.com/aziemp66/byte-bargain/common/jwt"
	passwordCommon "github.com/aziemp66/byte-bargain/common/password"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"

	userController "github.com/aziemp66/byte-bargain/internal/controller/user"
	userRepository "github.com/aziemp66/byte-bargain/internal/repository/user"
	userUseCase "github.com/aziemp66/byte-bargain/internal/usecase/user"

	productController "github.com/aziemp66/byte-bargain/internal/controller/product"
	productRepository "github.com/aziemp66/byte-bargain/internal/repository/product"
	productUseCase "github.com/aziemp66/byte-bargain/internal/usecase/product"

	webController "github.com/aziemp66/byte-bargain/internal/controller/web"
)

func main() {
	cfg := envCommon.LoadConfig()
	db := dbCommon.NewDB(cfg.DatabaseURL)

	goviewConfig := goview.Config{
		Root:         "web/views",
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
	passwordManager := passwordCommon.NewPasswordHashManager()
	jwtManager := jwtCommon.NewJWTManager(cfg.AccessTokenKey)

	mailDialer := gomail.NewDialer(cfg.EmailHost, cfg.EmailPort, cfg.EmailUsername, cfg.EmailPassword)

	httpServer.Router.Use(sessionManager.GetSessionHandler())

	httpServer.Router.Static("/product_image", "./public/product_image")
	httpServer.Router.Static("/static", "./web/static")

	api := httpServer.Router.Group("/api", httpCommon.MiddlewareErrorHandler(cfg.WebURL))

	UserRepository := userRepository.NewUserRepositoryImplementation()
	UserUseCase := userUseCase.NewUserUsecaseImplementation(UserRepository, db, passwordManager, jwtManager, mailDialer, cfg.WebURL)
	userController.NewUserController(api.Group("/user"), UserUseCase, sessionManager)

	ProductRepository := productRepository.NewProductRepositoryImplementation()
	ProductUseCase := productUseCase.NewProductUsecaseImplementation(ProductRepository, db, sessionManager)
	productController.NewProductController(api.Group("/product"), ProductUseCase)

	webController.NewWebController(httpServer.Router.Group(""), UserUseCase, ProductUseCase, sessionManager)

	err := httpServer.Router.Run(fmt.Sprintf(":%d", cfg.Port))

	if err != nil {
		panic(err)
	}
}
