package user

import (
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
)

type Usecase interface {
	Login(httpCommon.Login) error
	RegisterCustomer(httpCommon.RegisterCustomer) error
	RegisterSeller(httpCommon.RegisterSeller) error
	GetCustomerByID(customerID string) (httpCommon.Customer, error)
	GetSellerByID(sellerID string) (httpCommon.Seller, error)
	ForgotPassword(httpCommon.ForgotPassword) error
	ResetPassword(id, token string) error
	ChangePassword(id string, ChangePassword httpCommon.ChangePassword) error
	UpdateCustomerByID(customer httpCommon.Customer) error
	UpdateSellerByID(seller httpCommon.Seller) error
}
