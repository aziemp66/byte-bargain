package web

import (
	"github.com/gin-gonic/gin"

	productUC "github.com/aziemp66/byte-bargain/internal/usecase/product"
	userUC "github.com/aziemp66/byte-bargain/internal/usecase/user"
)

type WebController struct {
	UserUsecase    userUC.Usecase
	ProductUsecase productUC.Usecase
}

func NewWebController(router *gin.RouterGroup, userUsecase userUC.Usecase, productUsecase productUC.Usecase) {
	webController := &WebController{
		UserUsecase:    userUsecase,
		ProductUsecase: productUsecase,
	}

	//auth routes
	router.GET("/login", webController.Login)
	router.GET("/register", webController.Register)
	router.GET("/forgot-password", webController.ForgotPassword)
	router.GET("/reset-password/:id/:token", webController.ResetPassword)

	router.GET("/", webController.Home)
	router.GET("/product/:id", webController.ProductDetail)
	router.GET("/cart", webController.Cart)
	router.GET("/checkout", webController.Checkout)
	router.GET("/order", webController.Order)
	router.GET("/order/:id", webController.OrderDetail)
	router.GET("/profile", webController.Profile)
	router.GET("/profile/customer/:id", webController.CustomerProfile)
	router.GET("/profile/seller/:id", webController.SellerProfile)
}

func (w *WebController) Login(ctx *gin.Context) {

}

func (w *WebController) Register(ctx *gin.Context) {

}

func (w *WebController) ForgotPassword(ctx *gin.Context) {

}

func (w *WebController) ResetPassword(ctx *gin.Context) {

}

func (w *WebController) Home(ctx *gin.Context) {

}

func (w *WebController) ProductDetail(ctx *gin.Context) {

}

func (w *WebController) Cart(ctx *gin.Context) {

}

func (w *WebController) Checkout(ctx *gin.Context) {

}

func (w *WebController) Order(ctx *gin.Context) {

}

func (w *WebController) OrderDetail(ctx *gin.Context) {

}

func (w *WebController) Profile(ctx *gin.Context) {

}

func (w *WebController) CustomerProfile(ctx *gin.Context) {

}

func (w *WebController) SellerProfile(ctx *gin.Context) {

}
