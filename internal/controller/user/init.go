package user

import (
	userUseCase "github.com/aziemp66/byte-bargain/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase userUseCase.Usecase
}

func NewUserController(router *gin.RouterGroup, userUsecase userUseCase.Usecase) {
	userController := &UserController{
		UserUsecase: userUsecase,
	}

	router.POST("/login", userController.Login)
	router.POST("/register/customer", userController.RegisterCustomer)
	router.POST("/register/seller", userController.RegisterSeller)
	router.GET("/customer/:id", userController.GetCustomerByID)
	router.GET("/seller/:id", userController.GetSellerByID)
	router.POST("/forgot-password", userController.ForgotPassword)
	router.POST("/reset-password/:id/:token", userController.ResetPassword)
	router.POST("/change-password/:id", userController.ChangePassword)
	router.PUT("/customer/:id", userController.UpdateCustomerByID)
	router.PUT("/seller/:id", userController.UpdateSellerByID)

}

func (u *UserController) Login(c *gin.Context) {

}

func (u *UserController) RegisterCustomer(ctx *gin.Context) {

}

func (u *UserController) RegisterSeller(ctx *gin.Context) {

}

func (u *UserController) GetCustomerByID(ctx *gin.Context) {

}

func (u *UserController) GetSellerByID(ctx *gin.Context) {

}

func (u *UserController) ForgotPassword(ctx *gin.Context) {

}

func (u *UserController) ResetPassword(ctx *gin.Context) {

}

func (u *UserController) ChangePassword(ctx *gin.Context) {

}

func (u *UserController) UpdateCustomerByID(ctx *gin.Context) {

}

func (u *UserController) UpdateSellerByID(ctx *gin.Context) {

}

// Path: internal/controller/user/init.go
