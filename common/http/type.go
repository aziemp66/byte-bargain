package http

type (
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Value   interface{} `json:"value"`
	}

	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegisterCustomer struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		Name        string `json:"name"`
		Address     string `json:"address"`
		BirthDate   string `json:"birth_date"`
		PhoneNumber string `json:"phone_number"`
		Gender      string `json:"gender"`
	}

	RegisterSeller struct {
		Email          string `json:"email"`
		Password       string `json:"password"`
		Name           string `json:"name"`
		Address        string `json:"address"`
		BirthDate      string `json:"birth_date"`
		PhoneNumber    string `json:"phone_number"`
		Gender         string `json:"gender"`
		IdentityNumber string `json:"identity_number"`
		BankName       string `json:"bank_name"`
		DebitNumber    string `json:"debit_number"`
	}

	ForgotPassword struct {
		Email string `json:"email"`
	}

	ChangePassword struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	Customer struct {
		CustomerID  string `json:"customer_id"`
		UserID      string `json:"user_id"`
		Email       string `json:"email"`
		Name        string `json:"name"`
		Address     string `json:"address"`
		BirthDate   string `json:"birth_date"`
		PhoneNumber string `json:"phone_number"`
		Gender      string `json:"gender"`
	}

	Seller struct {
		SellerID       string `json:"seller_id"`
		UserID         string `json:"user_id"`
		Email          string `json:"email"`
		Name           string `json:"name"`
		Address        string `json:"address"`
		BirthDate      string `json:"birth_date"`
		PhoneNumber    string `json:"phone_number"`
		Gender         string `json:"gender"`
		IdentityNumber string `json:"identity_number"`
		BankName       string `json:"bank_name"`
		DebitNumber    string `json:"debit_number"`
	}

	Product struct {
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
		Category    string  `json:"category"`
		Description string  `json:"description"`
		Weight      float64 `json:"weight"`
	}

	productItem struct {
		ProductID string `json:"product_id"`
		Qty       int    `json:"qty"`
	}

	CreateOrder struct {
		SellerID string        `json:"seller_id"`
		Items    []productItem `json:"items"`
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
		Qty            int    `json:"qty"`
	}
	OrderItems struct {
		OrderID      string        `json:"order_id"`
		CustomerID   string        `json:"customer_id"`
		SellerID     string        `json:"seller_id"`
		OrderDate    string        `json:"order_date"`
		Status       string        `json:"status"`
		Products     []productItem `json:"products"`
		TotalPayment float64       `json:"total_payment"`
	}

	CartItems struct {
		CartID     string        `json:"cart_id"`
		CustomerID string        `json:"customer_id"`
		Products   []productItem `json:"products"`
	}
	CartProduct struct {
		CartProductID string `json:"cart_product_id"`
		CartID        string `json:"cart_id"`
		ProductID     string `json:"product_id"`
		Qty           int    `json:"qty"`
	}
	UpdateOrderStatus struct {
		OrderID string `json:"order_id"`
		Status  string `json:"status"`
	}

	UpdateOrderProductQty struct {
		OrderProductID string `json:"order_product_id"`
		Quantity       int    `json:"quantity"`
	}

	Payment struct {
		OrderID       string  `json:"order_id"`
		PaymentDate   string  `json:"payment_date"`
		TotalPayment  float64 `json:"total_payment"`
		PaymentMethod string  `json:"payment_method"`
	}
)
