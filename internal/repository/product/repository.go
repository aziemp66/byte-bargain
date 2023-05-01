package order

import (
	"context"
	"time"

	productDomain "github.com/aziemp66/byte-bargain/internal/domain/product"

	"database/sql"
)

type Repository interface {
	GetAllProduct(ctx context.Context, tx *sql.Tx) ([]productDomain.Product, error)
	GetSearchedProduct(ctx context.Context, tx *sql.Tx, search string) ([]productDomain.Product, error)
	GetAllProductBySellerID(ctx context.Context, tx *sql.Tx, sellerID string) ([]productDomain.Product, error)
	GetProductByID(ctx context.Context, tx *sql.Tx, productID string) (productDomain.Product, error)
	GetOrderByID(ctx context.Context, tx *sql.Tx, orderID string) (productDomain.Order, error)
	GetOrderByCustomerID(ctx context.Context, tx *sql.Tx, customerID string) ([]productDomain.Order, error)
	GetOrderBySellerID(ctx context.Context, tx *sql.Tx, sellerID string) ([]productDomain.Order, error)
	GetOrderProductByID(ctx context.Context, tx *sql.Tx, orderProductID string) (productDomain.OrderProduct, error)
	GetOrderProductByOrderID(ctx context.Context, tx *sql.Tx, orderID string) ([]productDomain.OrderProduct, error)
	GetCartProductByID(ctx context.Context, tx *sql.Tx, ID string) (productDomain.CartProduct, error)
	GetCartProductByCustomerID(ctx context.Context, tx *sql.Tx, customerID string) ([]productDomain.CartProduct, error)
	GetAllOrderProduct(ctx context.Context, tx *sql.Tx) ([]productDomain.OrderProduct, error)
	GetPaymentByID(ctx context.Context, tx *sql.Tx, paymentID string) (productDomain.Payment, error)
	InsertProduct(ctx context.Context, tx *sql.Tx, sellerID, productName string, price float64, stock int, category, description string, weight float64) error
	InsertOrder(ctx context.Context, tx *sql.Tx, customerID, sellerID string, orderDate time.Time, status string) error
	InsertOrderProduct(ctx context.Context, tx *sql.Tx, orderID, productID string, quantity int) error
	InsertCartProduct(ctx context.Context, tx *sql.Tx, customerID, productID string, quantity int) error
	InsertPayment(ctx context.Context, tx *sql.Tx, orderID string, paymentDate time.Time, totalPayment float64, paymentMethod string) error
	InsertImage(ctx context.Context, tx *sql.Tx, image string) (imageID string, err error)
	UpdateProductByID(ctx context.Context, tx *sql.Tx, productID, productName string, price float64, stock int, category, description string, weight float64) error
	UpdateOrderStatusByID(ctx context.Context, tx *sql.Tx, orderID, status string) error
	UpdateOrderProductQtyByID(ctx context.Context, tx *sql.Tx, orderProductID, quantity string) error
	UpdateCartProductQtyByID(ctx context.Context, tx *sql.Tx, cartProductID, quantity string) error
	UpdateLinkImageByID(ctx context.Context, tx *sql.Tx, imageID, productID string) error
	DeleteProductByID(ctx context.Context, tx *sql.Tx, productID string) error
	DeleteOrderProductByID(ctx context.Context, tx *sql.Tx, orderProductID string) error
	DeleteCartProductByID(ctx context.Context, tx *sql.Tx, cartProductID string) error
}
