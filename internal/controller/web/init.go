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
	router.GET("/register/customer", webController.RegisterCustomer)
	router.GET("/register/seller", webController.RegisterSeller)
	router.GET("/forgot-password", webController.ForgotPassword)
	router.GET("/reset-password/:id/:token", webController.ResetPassword)

	//non-auth routes
	router.GET("/", webController.Index)
	router.GET("/product/:id", webController.ProductDetail)
	router.GET("/profile/customer/:id", webController.CustomerProfile)
	router.GET("/profile/seller/:id", webController.SellerProfile)

	//customer routes
	customerRouter := router.Group("/customer")
	customerRouter.GET("/cart", webController.CustomerCart)
	customerRouter.GET("/checkout", webController.CustomerCheckout)
	customerRouter.GET("/order", webController.CustomerOrder)
	customerRouter.GET("/order/:id", webController.CustomerOrderDetail)
	customerRouter.GET("/profile", webController.CustomerSelfProfile)

	//seller routes
	sellerRouter := router.Group("/seller")
	sellerRouter.GET("/product", webController.SellerProduct)
	sellerRouter.GET("/product/add", webController.SellerProductAdd)
	sellerRouter.GET("/product/:id", webController.SellerProductDetail)
	sellerRouter.GET("/order", webController.SellerOrder)
	sellerRouter.GET("/order/:id", webController.SellerOrderDetail)
	sellerRouter.GET("/profile", webController.SellerSelfProfile)
}

func (w *WebController) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login", gin.H{})
}

func (w *WebController) RegisterCustomer(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "customer/register", gin.H{})
}

func (w *WebController) RegisterSeller(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "seller/register", gin.H{})
}

func (w *WebController) ForgotPassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "forgot-password", gin.H{})
}

func (w *WebController) ResetPassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "reset-password", gin.H{})
}

func (w *WebController) CustomerProfile(ctx *gin.Context) {
	customer, err := w.UserUsecase.GetCustomerByUserID(ctx, ctx.Param("id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	ctx.HTML(http.StatusOK, "customer-profile", gin.H{
		"customer": customer,
	})
}

func (w *WebController) SellerProfile(ctx *gin.Context) {
	seller, err := w.UserUsecase.GetSellerByUserID(ctx, ctx.Param("id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	ctx.HTML(http.StatusOK, "seller-profile", gin.H{
		"seller": seller,
	})
}

func (w *WebController) Index(ctx *gin.Context) {
	products, err := w.ProductUsecase.GetRecommendedProduct(ctx)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "index", gin.H{
		"products": products,
	})
}

func (w *WebController) ProductDetail(ctx *gin.Context) {
	product, err := w.ProductUsecase.GetProductByID(ctx, ctx.Param("id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "product-detail", gin.H{
		"product": product,
	})
}

func (w *WebController) CustomerCart(ctx *gin.Context) {
	cartItems, err := w.ProductUsecase.GetCustomerCart(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "cart", gin.H{
		"cart_items": cartItems,
	})
}

func (w *WebController) CustomerCheckout(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "checkout", gin.H{})
}

func (w *WebController) CustomerOrder(ctx *gin.Context) {
	customer, err := w.UserUsecase.GetCustomerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	orders, err := w.ProductUsecase.GetOrderByCustomerID(ctx, customer.CustomerID)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	ctx.HTML(http.StatusOK, "order", gin.H{
		"orders": orders,
	})
}

func (w *WebController) CustomerOrderDetail(ctx *gin.Context) {
	orderID := ctx.Param("id")

	order, err := w.ProductUsecase.GetOrderByID(ctx, orderID)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	customer, err := w.UserUsecase.GetCustomerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	if order.CustomerID != customer.CustomerID {
		ctx.HTML(http.StatusUnauthorized, "error", gin.H{
			"code":  "401",
			"error": "unauthorized",
		})

		return
	}

	ctx.HTML(http.StatusOK, "order-detail", gin.H{
		"order": order,
	})
}

func (w *WebController) CustomerSelfProfile(ctx *gin.Context) {
	customer, err := w.UserUsecase.GetCustomerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	ctx.HTML(http.StatusOK, "profile", gin.H{
		"customer": customer,
	})
}

func (w *WebController) SellerProduct(ctx *gin.Context) {
	seller, err := w.UserUsecase.GetSellerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	products, err := w.ProductUsecase.GetProductBySellerID(ctx, seller.SellerID)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	ctx.HTML(http.StatusOK, "seller-product", gin.H{
		"products": products,
	})
}

func (w *WebController) SellerProductAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "seller-product-add", gin.H{})
}

func (w *WebController) SellerProductDetail(ctx *gin.Context) {
	productID := ctx.Param("id")

	product, err := w.ProductUsecase.GetProductByID(ctx, productID)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	seller, err := w.UserUsecase.GetSellerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	if product.SellerID != seller.SellerID {
		ctx.HTML(http.StatusUnauthorized, "error", gin.H{
			"code":  "401",
			"error": "unauthorized",
		})

		return
	}

	ctx.HTML(http.StatusOK, "seller-product-detail", gin.H{
		"product": product,
	})
}

func (w *WebController) SellerOrder(ctx *gin.Context) {
	seller, err := w.UserUsecase.GetSellerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	orders, err := w.ProductUsecase.GetOrderBySellerID(ctx, seller.SellerID)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})
	}

	ctx.HTML(http.StatusOK, "seller-order", gin.H{
		"orders": orders,
	})
}

func (w *WebController) SellerOrderDetail(ctx *gin.Context) {
	orderID := ctx.Param("id")

	order, err := w.ProductUsecase.GetOrderByID(ctx, orderID)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	seller, err := w.UserUsecase.GetSellerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})
	}

	if order.SellerID != seller.SellerID {
		ctx.HTML(http.StatusUnauthorized, "error", gin.H{
			"code":  "401",
			"error": "unauthorized",
		})

		return
	}

	ctx.HTML(http.StatusOK, "seller-order-detail", gin.H{
		"order": order,
	})
}

func (w *WebController) SellerSelfProfile(ctx *gin.Context) {
	seller, err := w.UserUsecase.GetSellerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})

		return
	}

	ctx.HTML(http.StatusOK, "seller-self-profile", gin.H{
		"seller": seller,
	})
}
