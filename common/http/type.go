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

	UpdateOrderStatus struct {
		OrderID string `json:"order_id"`
		Status  string `json:"status"`
	}

	UpdateOrderProductQty struct {
		OrderProductID string `json:"order_product_id"`
		Quantity       int    `json:"quantity"`
	}
)
