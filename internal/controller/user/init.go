package user

import (
	"github.com/gin-gonic/gin"

	errorCommon "github.com/aziemp66/byte-bargain/common/error"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"

	userUseCase "github.com/aziemp66/byte-bargain/internal/usecase/user"
)

type UserController struct {
	UserUsecase    userUseCase.Usecase
	SessionManager *sessionCommon.SessionManager
}

func NewUserController(router *gin.RouterGroup, userUsecase userUseCase.Usecase, sessionManager *sessionCommon.SessionManager) {
	userController := &UserController{
		UserUsecase:    userUsecase,
		SessionManager: sessionManager,
	}

	router.POST("/login", userController.Login)
	router.POST("/register/customer", userController.RegisterCustomer)
	router.POST("/register/seller", userController.RegisterSeller)
	router.POST("/forgot-password", userController.ForgotPassword)
	router.POST("/reset-password", userController.ResetPassword)
	router.POST("/change-password/:id", userController.ChangePassword)
	router.PUT("/customer/:id", userController.UpdateCustomerByID)
	router.PUT("/seller/:id", userController.UpdateSellerByID)
}

func (u *UserController) Login(c *gin.Context) {
	var req httpCommon.Login

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	userID, err := u.UserUsecase.Login(c, req)

	if err != nil {
		c.Error(err)
		return
	}

	u.SessionManager.SetSessionValue(c, "user_id", userID)

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Login success",
	})
}

func (u *UserController) RegisterCustomer(ctx *gin.Context) {
	var req httpCommon.RegisterCustomer

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	err := u.UserUsecase.RegisterCustomer(ctx, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Customer register success",
	})
}

func (u *UserController) RegisterSeller(ctx *gin.Context) {
	var req httpCommon.RegisterSeller

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	err := u.UserUsecase.RegisterSeller(ctx, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Seller register success",
	})
}

func (u *UserController) ForgotPassword(ctx *gin.Context) {
	var req httpCommon.ForgotPassword

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	err := u.UserUsecase.ForgotPassword(ctx, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Forgot password request sent successfully",
	})
}

func (u *UserController) ResetPassword(ctx *gin.Context) {
	var req httpCommon.ResetPassword

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	err := u.UserUsecase.ResetPassword(ctx, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Reset password success",
	})
}

func (u *UserController) ChangePassword(ctx *gin.Context) {
	var req httpCommon.ChangePassword

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	userID := ctx.GetString("user_id")

	err := u.UserUsecase.ChangePassword(ctx, userID, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Change password success",
	})
}

func (u *UserController) UpdateCustomerByID(ctx *gin.Context) {
	var req httpCommon.UpdateCustomer

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	userID := ctx.GetString("user_id")

	err := u.UserUsecase.UpdateCustomerByID(ctx, userID, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Update customer success",
	})
}

func (u *UserController) UpdateSellerByID(ctx *gin.Context) {
	var req httpCommon.UpdateSeller

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	userID := ctx.GetString("user_id")

	err := u.UserUsecase.UpdateSellerByID(ctx, userID, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Update seller success",
	})
}

// Path: internal/controller/user/init.go
