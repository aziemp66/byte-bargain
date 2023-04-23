package order

import (
	"time"

	productDomain "github.com/aziemp66/byte-bargain/internal/domain/product"
	"github.com/gin-gonic/gin"

	"database/sql"
)

type Repository interface {
	GetAllProduct(ctx *gin.Context, tx *sql.Tx) ([]productDomain.Product, error)
	GetSearchedProduct(ctx *gin.Context, tx *sql.Tx, search string) ([]productDomain.Product, error)
	GetAllProductBySellerID(ctx *gin.Context, tx *sql.Tx, sellerID string) ([]productDomain.Product, error)
	GetProductByID(ctx *gin.Context, tx *sql.Tx, productID string) (productDomain.Product, error)
	GetOrderByID(ctx *gin.Context, tx *sql.Tx, orderID string) (productDomain.Order, error)
	GetOrderByCustomerID(ctx *gin.Context, tx *sql.Tx, customerID string) ([]productDomain.Order, error)
	GetOrderBySellerID(ctx *gin.Context, tx *sql.Tx, sellerID string) ([]productDomain.Order, error)
	GetOrderProductByID(ctx *gin.Context, tx *sql.Tx, orderProductID string) (productDomain.OrderProduct, error)
	GetOrderProductByOrderID(ctx *gin.Context, tx *sql.Tx, orderID string) ([]productDomain.OrderProduct, error)
	GetCartProductByCustomerID(ctx *gin.Context, tx *sql.Tx, customerID string) ([]productDomain.CartProduct, error)
	GetAllOrderProduct(ctx *gin.Context, tx *sql.Tx) ([]productDomain.OrderProduct, error)
	GetPaymentByID(ctx *gin.Context, tx *sql.Tx, paymentID string) (productDomain.Payment, error)
	InsertProduct(ctx *gin.Context, tx *sql.Tx, sellerID, productName string, price float64, stock int, category, description string, weight float64) error
	InsertOrder(ctx *gin.Context, tx *sql.Tx, customerID, sellerID string, orderDate time.Time, status string) error
	InsertOrderProduct(ctx *gin.Context, tx *sql.Tx, orderID, productID string, quantity int) error
	InsertCartProduct(ctx *gin.Context, tx *sql.Tx, cartID, productID string, quantity int) error
	InsertPayment(ctx *gin.Context, tx *sql.Tx, orderID string, paymentDate time.Time, totalPayment float64, paymentMethod string) error
	InsertImage(ctx *gin.Context, tx *sql.Tx, image string) (imageID string, err error)
	UpdateProductByID(ctx *gin.Context, tx *sql.Tx, productID, productName string, price float64, stock int, category, description string, weight float64) error
	UpdateOrderStatusByID(ctx *gin.Context, tx *sql.Tx, orderID, status string) error
	UpdateOrderProductQtyByID(ctx *gin.Context, tx *sql.Tx, orderProductID, quantity string) error
	UpdateCartProductQtyByID(ctx *gin.Context, tx *sql.Tx, cartProductID, quantity string) error
	UpdateLinkImageByID(ctx *gin.Context, tx *sql.Tx, imageID, productID string) error
	DeleteProductByID(ctx *gin.Context, tx *sql.Tx, productID string) error
	DeleteOrderProductByID(ctx *gin.Context, tx *sql.Tx, orderProductID string) error
	DeleteCartProductByID(ctx *gin.Context, tx *sql.Tx, cartProductID string) error
}
