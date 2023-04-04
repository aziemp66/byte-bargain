package order

import (
	productDomain "github.com/aziemp66/byte-bargain/internal/domain/product"

	"context"
	"database/sql"
)

type Repository interface {
	GetRecommendedProduct(ctx context.Context, tx *sql.Tx) ([]productDomain.Product, error)
	GetSearchedProduct(ctx context.Context, tx *sql.Tx, search string) ([]productDomain.Product, error)
	GetAllProductBySellerID(ctx context.Context, tx *sql.Tx, sellerID string) ([]productDomain.Product, error)
	GetProductByID(ctx context.Context, tx *sql.Tx, productID string) (productDomain.Product, error)
	GetOrderByID(ctx context.Context, tx *sql.Tx, orderID string) (productDomain.Order, error)
	GetOrderProductByID(ctx context.Context, tx *sql.Tx, orderProductID string) (productDomain.OrderProduct, error)
	GetPaymentByID(ctx context.Context, tx *sql.Tx, paymentID string) (productDomain.Payment, error)
	InsertProduct(ctx context.Context, tx *sql.Tx, sellerID, productName, price, stock, category, description, weight string) error
	InsertOrder(ctx context.Context, tx *sql.Tx, customerID, sellerID, orderDate, status string) error
	InsertOrderProduct(ctx context.Context, tx *sql.Tx, orderID, productID, quantity string) error
	InsertPayment(ctx context.Context, tx *sql.Tx, orderID, paymentDate, totalPayment, paymentMethod string) error
	UpdateProductByID(ctx context.Context, tx *sql.Tx, productID, sellerID, productName, price, stock, category, description, weight string) error
	UpdateOrderStatusByID(ctx context.Context, tx *sql.Tx, orderID, status string) error
	UpdateOrderProductQtyByID(ctx context.Context, tx *sql.Tx, orderProductID, quantity string) error
	DeleteProductByID(ctx context.Context, tx *sql.Tx, productID string) error
	DeleteOrderProductByID(ctx context.Context, tx *sql.Tx, orderProductID string) error
}
