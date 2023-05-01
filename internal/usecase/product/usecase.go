package product

import (
	"context"

	httpCommon "github.com/aziemp66/byte-bargain/common/http"
)

type Usecase interface {
	GetRecommendedProduct(ctx context.Context) ([]httpCommon.Product, error)
	GetSearchedProduct(ctx context.Context, search string) ([]httpCommon.Product, error)
	GetProductBySellerID(ctx context.Context, sellerID string) ([]httpCommon.Product, error)
	GetProductByID(ctx context.Context, productID string) (httpCommon.Product, error)
	GetOrderByID(ctx context.Context, orderID string) (httpCommon.Order, error)
	GetOrderByCustomerID(ctx context.Context, customerID string) ([]httpCommon.OrderItems, error)
	GetOrderBySellerID(ctx context.Context, sellerID string) ([]httpCommon.OrderItems, error)
	GetOrderProductByID(ctx context.Context, orderProductID string) (httpCommon.OrderProduct, error)
	GetCustomerCart(ctx context.Context, customerID string) (httpCommon.Cart, error)
	GetCartProductByID(ctx context.Context, cartProductID string) (httpCommon.CartProduct, error)
	GetPaymentByID(ctx context.Context, paymentID string) (httpCommon.Payment, error)
	InsertProduct(ctx context.Context, sellerID string, product httpCommon.ProductRequest) error
	InsertOrder(ctx context.Context, CustomerID string, createOrder httpCommon.CreateOrder) error
	InsertOrderProduct(ctx context.Context, orderProduct httpCommon.OrderProduct) error
	InsertCartProduct(ctx context.Context, customerID string, cartProduct httpCommon.AddCartProduct) error
	InsertPayment(ctx context.Context, payment httpCommon.Payment) error
	InsertImage(ctx context.Context, filepath string) (imagesID string, err error)
	UpdateProductByID(ctx context.Context, productID string, product httpCommon.ProductRequest) error
	UpdateOrderStatusByID(ctx context.Context, updateOrderStatus httpCommon.UpdateOrderStatus) error
	UpdateOrderProductQtyByID(ctx context.Context, orderProductID, quantity string) error
	UpdateCartProductQtyByID(ctx context.Context, cartProductID, quantity string) error
	DeleteProductByID(ctx context.Context, productID string) error
	DeleteCartProductByID(ctx context.Context, cartProductID string) error
}
