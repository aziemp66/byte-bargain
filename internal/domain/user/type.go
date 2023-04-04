package user

import "time"

type (
	User struct {
		UserID   string `json:"user_id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	Customer struct {
		CustomerID  string    `json:"customer_id"`
		UserID      string    `json:"user_id"`
		Name        string    `json:"name,omitempty"`
		Address     string    `json:"address,omitempty"`
		BirthDate   time.Time `json:"birth_date,omitempty"`
		PhoneNumber string    `json:"phone_number,omitempty"`
		Gender      string    `json:"gender,omitempty"`
	}
	Seller struct {
		SellerID       string    `json:"seller_id"`
		UserID         string    `json:"user_id"`
		Name           string    `json:"name,omitempty"`
		Address        string    `json:"address"`
		BirthDate      time.Time `json:"birth_date"`
		PhoneNumber    string    `json:"phone_number"`
		Gender         string    `json:"gender"`
		IdentityNumber string    `json:"identity_number"`
		BankName       string    `json:"bank_name"`
		DebitNumber    string    `json:"debit_number"`
	}
)
