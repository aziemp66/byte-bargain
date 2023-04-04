package transcation

import "time"

type (
	Product struct {
		ProductID   string  `json:"product_id"`
		SellerID    string  `json:"seller_id"`
		ProductName string  `json:"product_name"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
		Category    string  `json:"category"`
		Description string  `json:"description"`
		Weight      float64 `json:"weight"`
	}
	Order struct {
		OrderID    string `json:"order_id"`
		CustomerID string `json:"customer_id"`
		SellerID   string `json:"seller_id"`
		OrderDate  string `json:"order_date"`
		Status     string `json:"status"`
	}
	OrderProduct struct {
		OrderProductID string `json:"order_product_id"`
		OrderID        string `json:"order_id"`
		ProductID      string `json:"product_id"`
		Quantity       int    `json:"quantity"`
	}
	Payment struct {
		PaymentID     string    `json:"payment_id"`
		OrderID       string    `json:"order_id"`
		PaymentDate   time.Time `json:"payment_date"`
		TotalPayment  float64   `json:"total_payment"`
		PaymentMethod string    `json:"payment_method"`
	}
)
