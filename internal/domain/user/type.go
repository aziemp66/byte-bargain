package user

import "time"

type (
	User struct {
		UserID   string
		Email    string
		Password string
	}
	Customer struct {
		CustomerID  string
		UserID      string
		Name        string
		Address     string
		BirthDate   time.Time
		PhoneNumber string
		Gender      string
	}
	Seller struct {
		SellerID       string
		UserID         string
		Name           string
		Address        string
		BirthDate      time.Time
		PhoneNumber    string
		Gender         string
		IdentityNumber string
		BankName       string
		DebitNumber    string
	}
)
