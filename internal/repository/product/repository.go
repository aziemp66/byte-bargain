package order

import (
	productDomain "github.com/aziemp66/byte-bargain/internal/domain/product"
)

type Repository interface {
	GetRecommendedProduct() ([]productDomain.Product, error)
	GetSearchedProduct(search string) ([]productDomain.Product, error)
	GetAllProductBySellerID(sellerID string) ([]productDomain.Product, error)
	GetProductByID(productID string) (productDomain.Product, error)
	GetOrderByID(orderID string) (productDomain.Order, error)
	GetOrderProductByID(orderProductID string) (productDomain.OrderProduct, error)
	GetPaymentByID(paymentID string) (productDomain.Payment, error)
	InsertProduct(sellerID, productName, price, stock, category, description, weight string) error
	InsertOrder(customerID, sellerID, orderDate, status string) error
	InsertOrderProduct(orderID, productID, quantity string) error
	InsertPayment(orderID, paymentDate, totalPayment, paymentMethod string) error
	UpdateProductByID(productID, sellerID, productName, price, stock, category, description, weight string) error
	UpdateOrderStatusByID(orderID, status string) error
	UpdateOrderProductQtyByID(orderProductID, quantity string) error
	DeleteProductByID(productID string) error
	DeleteOrderProductByID(orderProductID string) error
}
