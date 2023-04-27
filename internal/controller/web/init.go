package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	sessionCommon "github.com/aziemp66/byte-bargain/common/session"

	productUC "github.com/aziemp66/byte-bargain/internal/usecase/product"
	userUC "github.com/aziemp66/byte-bargain/internal/usecase/user"
)

type WebController struct {
	UserUsecase    userUC.Usecase
	ProductUsecase productUC.Usecase
	SessionManager *sessionCommon.SessionManager
}

func NewWebController(router *gin.RouterGroup, userUsecase userUC.Usecase, productUsecase productUC.Usecase, sessionManager *sessionCommon.SessionManager) {
	webController := &WebController{
		UserUsecase:    userUsecase,
		ProductUsecase: productUsecase,
		SessionManager: sessionManager,
	}

	//auth routes
	router.GET("/login", webController.Login)
	router.GET("/register", webController.Register)
	router.GET("/forgot-password", webController.ForgotPassword)
	router.GET("/reset-password/:id/:token", webController.ResetPassword)

	router.GET("/", webController.Index)
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
	ctx.HTML(http.StatusOK, "login", gin.H{})
}

func (w *WebController) Register(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register", gin.H{})
}

func (w *WebController) ForgotPassword(ctx *gin.Context) {

}

func (w *WebController) ResetPassword(ctx *gin.Context) {

}

func (w *WebController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", gin.H{})
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
