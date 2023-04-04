package transcation

import "time"

type (
	Product struct {
		ProductID   string
		SellerID    string
		Name        string
		Price       float64
		Stock       int
		Category    string
		Description string
		Weight      float64
	}
	Order struct {
		OrderID    string
		CustomerID string
		SellerID   string
		OrderDate  string
		Status     string
	}
	OrderProduct struct {
		OrderProductID string
		OrderID        string
		ProductID      string
		Quantity       int
	}
	Payment struct {
		PaymentID     string    `json:"payment_id"`
		OrderID       string    `json:"order_id"`
		PaymentDate   time.Time `json:"payment_date"`
		TotalPayment  float64   `json:"total_payment"`
		PaymentMethod string    `json:"payment_method"`
	}
)
