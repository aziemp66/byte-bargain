package user

import (
	userDomain "github.com/aziemp66/byte-bargain/internal/domain/user"
)

type Repository interface {
	GetUserByID(userID string) (userDomain.User, error)
	GetUserByEmail(email string) (userDomain.User, error)
	GetCustomerByID(customerID string) (userDomain.Customer, error)
	GetSellerByID(sellerID string) (userDomain.Seller, error)
	InsertUser(email, password string) error
	InsertCustomer(userID, name, address, phoneNumber, gender string) error
	InsertSeller(userID, name, address, birthdate, phoneNumber, gender, identityNumber, bankName, debitNumber string) error
	UpdateCustomerByID(userID, name, address, phoneNumber, gender string) error
	UpdateSellerByID(userID, name, address, birthdate, phoneNumber, gender, identityNumber, bankName, debitNumber string) error
	UpdateUserPasswordByID(userID string, password string) error
}
