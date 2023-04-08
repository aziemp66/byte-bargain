package user

import (
	userUseCase "github.com/aziemp66/byte-bargain/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase userUseCase.Usecase
}

func NewUserController(userUsecase userUseCase.Usecase) *UserController {
	return &UserController{
		UserUsecase: userUsecase,
	}
}

func (u *UserController) Login(ctx *gin.Context) error {
	return nil
}

func (u *UserController) RegisterCustomer(ctx *gin.Context) error {
	return nil
}

func (u *UserController) RegisterSeller(ctx *gin.Context) error {
	return nil
}

func (u *UserController) GetCustomerByID(ctx *gin.Context) error {
	return nil
}

func (u *UserController) GetSellerByID(ctx *gin.Context) error {
	return nil
}

func (u *UserController) ForgotPassword(ctx *gin.Context) error {
	return nil
}

func (u *UserController) ResetPassword(ctx *gin.Context) error {
	return nil
}

func (u *UserController) ChangePassword(ctx *gin.Context) error {
	return nil
}

func (u *UserController) UpdateCustomerByID(ctx *gin.Context) error {
	return nil
}

func (u *UserController) UpdateSellerByID(ctx *gin.Context) error {
	return nil
}

// Path: internal/controller/user/init.go
