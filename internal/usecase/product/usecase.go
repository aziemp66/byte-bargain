package product

import (
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	"github.com/gin-gonic/gin"
)

type Usecase interface {
	GetRecommendedProduct(ctx *gin.Context) ([]httpCommon.Product, error)
	GetSearchedProduct(ctx *gin.Context, search string) ([]httpCommon.Product, error)
	GetProductBySellerID(ctx *gin.Context, sellerID string) ([]httpCommon.Product, error)
	GetProductByID(ctx *gin.Context, productID string) (httpCommon.Product, error)
	GetOrderByID(ctx *gin.Context, orderID string) (httpCommon.Order, error)
	GetOrderByCustomerID(ctx *gin.Context, customerID string) ([]httpCommon.OrderItems, error)
	GetSellerOrderByID(ctx *gin.Context, sellerID string) ([]httpCommon.OrderItems, error)
	GetOrderProductByID(ctx *gin.Context, orderProductID string) (httpCommon.OrderProduct, error)
	GetCustomerCart(ctx *gin.Context, customerID string) (httpCommon.Cart, error)
	GetPaymentByID(ctx *gin.Context, paymentID string) (httpCommon.Payment, error)
	InsertProduct(ctx *gin.Context, product httpCommon.AddProduct) error
	InsertOrder(ctx *gin.Context, createOrder httpCommon.CreateOrder) error
	InsertOrderProduct(ctx *gin.Context, orderProduct httpCommon.OrderProduct) error
	InsertCartProduct(ctx *gin.Context, cartProduct httpCommon.CartProduct) error
	InsertPayment(ctx *gin.Context, payment httpCommon.Payment) error
	InsertImages(ctx *gin.Context) (imagesID []string, err error)
	UpdateProductByID(ctx *gin.Context, productID string, product httpCommon.Product) error
	UpdateOrderStatusByID(ctx *gin.Context, orderID, status string) error
	UpdateOrderProductQtyByID(ctx *gin.Context, orderProductID, quantity string) error
	UpdateCartProductQtyByID(ctx *gin.Context, cartProductID, quantity string) error
	DeleteProductByID(ctx *gin.Context, productID string) error
	DeleteOrderProductByID(ctx *gin.Context, orderProductID string) error
	DeleteCartProductByID(ctx *gin.Context, cartProductID string) error
}
