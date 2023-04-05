package product

import (
	httpCommon "github.com/aziemp66/byte-bargain/common/http"

	"context"
)

type Usecase interface {
	GetRecommendedProduct(ctx context.Context) ([]httpCommon.Product, error)
	GetSearchedProduct(ctx context.Context, search string) ([]httpCommon.Product, error)
	GetAllProductBySellerID(ctx context.Context, sellerID string) ([]httpCommon.Product, error)
	GetProductByID(ctx context.Context, productID string) (httpCommon.Product, error)
	GetOrderByID(ctx context.Context, orderID string) (httpCommon.Order, error)
	GetCustomerOrderByID(ctx context.Context, customerID string) ([]httpCommon.OrderItems, error)
	GetSellerOrderByID(ctx context.Context, sellerID string) ([]httpCommon.OrderItems, error)
	GetOrderProductByID(ctx context.Context, orderProductID string) (httpCommon.OrderProduct, error)
	GetCustomerCart(ctx context.Context, customerID string) (httpCommon.CartItems, error)
	GetPaymentByID(ctx context.Context, paymentID string) (httpCommon.Payment, error)
	InsertProduct(ctx context.Context, product httpCommon.Product) error
	InsertOrder(ctx context.Context, createOrder httpCommon.CreateOrder) error
	InsertOrderProduct(ctx context.Context, orderProduct httpCommon.OrderProduct) error
	InsertCartProduct(ctx context.Context, cartProduct httpCommon.CartProduct) error
	InsertPayment(ctx context.Context, payment httpCommon.Payment) error
	UpdateProductByID(ctx context.Context, productID string, product httpCommon.Product) error
	UpdateOrderStatusByID(ctx context.Context, orderID, status string) error
	UpdateOrderProductQtyByID(ctx context.Context, orderProductID, quantity string) error
	UpdateCartProductQtyByID(ctx context.Context, cartProductID, quantity string) error
	DeleteProductByID(ctx context.Context, productID string) error
	DeleteOrderProductByID(ctx context.Context, orderProductID string) error
}
