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
	GetOrderProductByID(orderProductID string) (httpCommon.OrderProduct, error)
	GetPaymentByID(paymentID string) (httpCommon.Payment, error)
	InsertProduct(httpCommon.Product) error
	InsertOrder(httpCommon.CreateOrder) error
	InsertOrderProduct(httpCommon.OrderProduct) error
	InsertPayment(httpCommon.Payment) error
	UpdateProductByID(productID string, product httpCommon.Product) error
	UpdateOrderStatusByID(orderID, status string) error
	UpdateOrderProductQtyByID(orderProductID, quantity string) error
	DeleteProductByID(productID string) error
	DeleteOrderProductByID(orderProductID string) error
}
