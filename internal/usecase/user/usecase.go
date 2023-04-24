package user

import (
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	"github.com/gin-gonic/gin"
)

type Usecase interface {
	Login(ctx *gin.Context, login httpCommon.Login) error
	RegisterCustomer(ctx *gin.Context, registerCustomer httpCommon.RegisterCustomer) error
	RegisterSeller(ctx *gin.Context, registerSeller httpCommon.RegisterSeller) error
	GetCustomerByID(ctx *gin.Context, customerID string) (httpCommon.Customer, error)
	GetSellerByID(ctx *gin.Context, sellerID string) (httpCommon.Seller, error)
	ForgotPassword(ctx *gin.Context, forgotPassword httpCommon.ForgotPassword) error
	ResetPassword(ctx *gin.Context, resetPassword httpCommon.ResetPassword) error
	ChangePassword(ctx *gin.Context, ChangePassword httpCommon.ChangePassword) error
	UpdateCustomerByID(ctx *gin.Context, customer httpCommon.UpdateCustomer) error
	UpdateSellerByID(ctx *gin.Context, seller httpCommon.UpdateSeller) error
	ActivateAccount(ctx *gin.Context, token string) error
	SendActivationEmail(ctx *gin.Context, email string) error
}
