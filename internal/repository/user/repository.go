package user

import (
	UserDomain "github.com/aziemp66/byte-bargain/internal/domain/user"
)

type Repository interface {
	GetUserByID(userID string) (UserDomain.User, error)
	GetUserByEmail(email string) (UserDomain.User, error)
	GetCustomerByID(customerID string) (UserDomain.Customer, error)
	GetSellerByID(sellerID string) (UserDomain.Seller, error)
	InsertUser(user UserDomain.User) error
	InsertCustomer(customer UserDomain.Customer) error
	InsertSeller(seller UserDomain.Seller) error
	UpdateCustomer(customer UserDomain.Customer) error
	UpdateSeller(seller UserDomain.Seller) error
	UpdateUserPassword(userID string, password string) error
}
