package transcation

import "time"

type (
	Product struct {
		ProductID   string  `json:"product_id"`
		SellerID    string  `json:"seller_id"`
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
		Category    string  `json:"category"`
		Description string  `json:"description"`
		Weight      float64 `json:"weight"`
	}
	ProductImage struct {
		ProductImageID string `json:"product_image_id"`
		ProductID      string `json:"product_id"`
		Image          string `json:"image"`
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
	CartProduct struct {
		CartProductID string `json:"cart_product_id"`
		CustomerID    string `json:"customer_id"`
		ProductID     string `json:"product_id"`
		Quantity      int    `json:"quantity"`
	}
	Payment struct {
		PaymentID     string    `json:"payment_id"`
		OrderID       string    `json:"order_id"`
		PaymentDate   time.Time `json:"payment_date"`
		TotalPayment  float64   `json:"total_payment"`
		PaymentMethod string    `json:"payment_method"`
	}
)
