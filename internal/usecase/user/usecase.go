package user

import (
	"context"

	httpCommon "github.com/aziemp66/byte-bargain/common/http"
)

type Usecase interface {
	Login(ctx context.Context, login httpCommon.Login) (userID string, err error)
	RegisterCustomer(ctx context.Context, registerCustomer httpCommon.RegisterCustomer) error
	RegisterSeller(ctx context.Context, registerSeller httpCommon.RegisterSeller) error
	GetCustomerByUserID(ctx context.Context, userID string) (httpCommon.Customer, error)
	GetSellerByUserID(ctx context.Context, userID string) (httpCommon.Seller, error)
	ForgotPassword(ctx context.Context, forgotPassword httpCommon.ForgotPassword) error
	ResetPassword(ctx context.Context, resetPassword httpCommon.ResetPassword) error
	ChangePassword(ctx context.Context, userID string, ChangePassword httpCommon.ChangePassword) error
	UpdateCustomerByID(ctx context.Context, customerID string, customer httpCommon.UpdateCustomer) error
	UpdateSellerByID(ctx context.Context, sellerID string, seller httpCommon.UpdateSeller) error
	ActivateAccount(ctx context.Context, token string) error
	SendActivationEmail(ctx context.Context, email string) error
}
