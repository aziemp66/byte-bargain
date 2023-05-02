package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	sessionCommon "github.com/aziemp66/byte-bargain/common/session"

	productUC "github.com/aziemp66/byte-bargain/internal/usecase/product"
	userUC "github.com/aziemp66/byte-bargain/internal/usecase/user"
)

type WebView struct {
	UserUsecase    userUC.Usecase
	ProductUsecase productUC.Usecase
	SessionManager *sessionCommon.SessionManager
}

func NewWebView(router *gin.RouterGroup, userUsecase userUC.Usecase, productUsecase productUC.Usecase, sessionManager *sessionCommon.SessionManager) {
	webController := &WebView{
		UserUsecase:    userUsecase,
		ProductUsecase: productUsecase,
		SessionManager: sessionManager,
	}

	router.GET("/login", webController.Login)//done
	router.GET("/register/customer", webController.RegisterCustomer)//done
	router.GET("/register/seller", webController.RegisterSeller)//done
	router.GET("/forgot-password", webController.ForgotPassword)//done
	router.GET("/reset-password", webController.ResetPassword)//done

	//non-auth routes
	router.GET("/", webController.Index)//done
	router.GET("/product/:id", webController.ProductDetail)
	router.GET("/product/seller/:sellerID", webController.ProductBySeller)
	router.GET("/profile/customer/:id", webController.CustomerProfile)
	router.GET("/profile/seller/:id", webController.SellerProfile)

	//customer-auth routes
	customerRouter := router.Group("/customer")
	customerRouter.GET("/cart", webController.CustomerCart)
	customerRouter.GET("/checkout", webController.CustomerCheckout)
	customerRouter.GET("/order", webController.CustomerOrder)
	customerRouter.GET("/order/:id", webController.CustomerOrderDetail)
	customerRouter.GET("/profile", webController.CustomerSelfProfile)

	//seller-auth routes
	sellerRouter := router.Group("/seller")
	sellerRouter.GET("/product", webController.SellerProduct)
	sellerRouter.GET("/product/add", webController.SellerProductAdd)
	sellerRouter.GET("/product/:id", webController.SellerProductDetail)
	sellerRouter.GET("/order", webController.SellerOrder)
	sellerRouter.GET("/order/:id", webController.SellerOrderDetail)
	sellerRouter.GET("/profile", webController.SellerSelfProfile)
}

func (w *WebView) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login", gin.H{})
}

func (w *WebView) RegisterCustomer(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "customer/register", gin.H{})
}

func (w *WebView) RegisterSeller(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "seller/register", gin.H{})
}

func (w *WebView) ForgotPassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "forgot-password", gin.H{})
}

func (w *WebView) ResetPassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "reset-password", gin.H{})
}

func (w *WebView) CustomerProfile(ctx *gin.Context) {
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

func (w *WebView) SellerProfile(ctx *gin.Context) {
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

func (w *WebView) Index(ctx *gin.Context) {
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

func (w *WebView) ProductDetail(ctx *gin.Context) {
	// product, err := w.ProductUsecase.GetProductByID(ctx, ctx.Param("id"))

	// if err != nil {
	// 	ctx.HTML(http.StatusInternalServerError, "error", gin.H{
	// 		"code":  "500",
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// seller, err := w.UserUsecase.GetSellerByUserID(ctx, product.SellerID)

	// if err != nil {
	// 	ctx.HTML(http.StatusInternalServerError, "error", gin.H{
	// 		"code":  "500",
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// product.SellerName = seller.Name

	ctx.HTML(http.StatusOK, "product-detail", gin.H{
		// "product": product,
	})
}

func (w *WebView) ProductBySeller(ctx *gin.Context) {
	products, err := w.ProductUsecase.GetProductBySellerID(ctx, ctx.Param("sellerID"))

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "product-by-seller", gin.H{
		"products": products,
	})
}

func (w *WebView) CustomerCart(ctx *gin.Context) {
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

func (w *WebView) CustomerCheckout(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "checkout", gin.H{})
}

func (w *WebView) CustomerOrder(ctx *gin.Context) {
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

func (w *WebView) CustomerOrderDetail(ctx *gin.Context) {
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

func (w *WebView) CustomerSelfProfile(ctx *gin.Context) {
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

func (w *WebView) SellerProduct(ctx *gin.Context) {
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

func (w *WebView) SellerProductAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "seller-product-add", gin.H{})
}

func (w *WebView) SellerProductDetail(ctx *gin.Context) {
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

func (w *WebView) SellerOrder(ctx *gin.Context) {
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

func (w *WebView) SellerOrderDetail(ctx *gin.Context) {
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

func (w *WebView) SellerSelfProfile(ctx *gin.Context) {
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
