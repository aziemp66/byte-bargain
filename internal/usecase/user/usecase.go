package user

import (
	httpCommon "github.com/aziemp66/byte-bargain/common/http"

	"context"
)

type Usecase interface {
	Login(ctx context.Context, login httpCommon.Login) error
	RegisterCustomer(ctx context.Context, registerCustomer httpCommon.RegisterCustomer) error
	RegisterSeller(ctx context.Context, registerSeller httpCommon.RegisterSeller) error
	GetCustomerByID(ctx context.Context, customerID string) (httpCommon.Customer, error)
	GetSellerByID(ctx context.Context, sellerID string) (httpCommon.Seller, error)
	ForgotPassword(ctx context.Context, forgotPassword httpCommon.ForgotPassword) error
	ResetPassword(ctx context.Context, id, token string) error
	ChangePassword(ctx context.Context, id string, ChangePassword httpCommon.ChangePassword) error
	UpdateCustomerByID(ctx context.Context, customer httpCommon.Customer) error
	UpdateSellerByID(ctx context.Context, seller httpCommon.Seller) error
}
