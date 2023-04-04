package order

import (
	orderDomain "github.com/aziemp66/byte-bargain/internal/domain/order"
)

type Repository interface {
	GetProductByID(productID string) (orderDomain.Product, error)
	GetOrderByID(orderID string) (orderDomain.Order, error)
	GetOrderProductByID(orderProductID string) (orderDomain.OrderProduct, error)
	GetPaymentByID(paymentID string) (orderDomain.Payment, error)
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
