package product

import (
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
)

type Usecase interface {
	GetRecommendedProduct() ([]httpCommon.Product, error)
	GetSearchedProduct(search string) ([]httpCommon.Product, error)
	GetAllProductBySellerID(sellerID string) ([]httpCommon.Product, error)
	GetProductByID(productID string) (httpCommon.Product, error)
	GetOrderByID(orderID string) (httpCommon.Order, error)
	GetCustomerOrderByID(customerID string) ([]httpCommon.OrderItems, error)
	GetSellerOrderByID(sellerID string) ([]httpCommon.OrderItems, error)
	GetOrderProductByID(orderProductID string) (httpCommon.OrderProduct, error)
	GetCustomerCart(customerID string) (httpCommon.CartItems, error)
	GetPaymentByID(paymentID string) (httpCommon.Payment, error)
	InsertProduct(httpCommon.Product) error
	InsertOrder(httpCommon.CreateOrder) error
	InsertOrderProduct(httpCommon.OrderProduct) error
	InsertCartProduct(httpCommon.CartProduct) error
	InsertPayment(httpCommon.Payment) error
	UpdateProductByID(productID string, product httpCommon.Product) error
	UpdateOrderStatusByID(orderID, status string) error
	UpdateOrderProductQtyByID(orderProductID, quantity string) error
	UpdateCartProductQtyByID(cartProductID, quantity string) error
	DeleteProductByID(productID string) error
	DeleteOrderProductByID(orderProductID string) error
}
