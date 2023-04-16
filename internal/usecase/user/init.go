package user

import (
	"database/sql"
	"time"

	dbCommon "github.com/aziemp66/byte-bargain/common/db"
	errorCommon "github.com/aziemp66/byte-bargain/common/error"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	passwordCommon "github.com/aziemp66/byte-bargain/common/password"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"
	"github.com/gin-gonic/gin"

	userRepository "github.com/aziemp66/byte-bargain/internal/repository/user"
)

type UserUsecaseImplementation struct {
	UserRepository      userRepository.Repository
	DB                  *sql.DB
	SessionManager      *sessionCommon.SessionManager
	PasswordHashManager *passwordCommon.PasswordHashManager
}

func NewUserUsecaseImplementation(userRepository userRepository.Repository, db *sql.DB, sessionManager *sessionCommon.SessionManager, passwordManager *passwordCommon.PasswordHashManager) *UserUsecaseImplementation {
	return &UserUsecaseImplementation{
		UserRepository:      userRepository,
		DB:                  db,
		SessionManager:      sessionManager,
		PasswordHashManager: passwordManager,
	}
}

func (u *UserUsecaseImplementation) Login(ctx *gin.Context, login httpCommon.Login) error {
	tx, err := u.DB.Begin()

	if err != nil {
		return errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	user, err := u.UserRepository.GetUserByEmail(ctx, tx, login.Email)

	if err != nil {
		return err
	}

	err = u.PasswordHashManager.CheckPasswordHash(login.Password, user.Password)

	if err != nil {
		return errorCommon.NewInvariantError("invalid password")
	}

	err = u.SessionManager.SetSessionValue(ctx, "user_id", user.UserID)

	if err != nil {
		return errorCommon.NewInvariantError("failed to set session value")
	}

	return nil
}

func (u *UserUsecaseImplementation) RegisterCustomer(ctx *gin.Context, registerCustomer httpCommon.RegisterCustomer) error {
	tx, err := u.DB.Begin()

	if err != nil {
		return errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	user, err := u.UserRepository.GetUserByEmail(ctx, tx, registerCustomer.Email)

	if err != nil {
		return err
	}

	if user.UserID != "" {
		return errorCommon.NewInvariantError("email already registered")
	}

	hashedPassword, err := u.PasswordHashManager.HashPassword(registerCustomer.Password)

	if err != nil {
		return errorCommon.NewInvariantError("failed to hash password")
	}

	userId, err := u.UserRepository.InsertUser(ctx, tx, registerCustomer.Email, hashedPassword)

	if err != nil {
		return err
	}

	userBirthdate, err := time.Parse("2006-01-02", registerCustomer.BirthDate)

	if err != nil {
		return errorCommon.NewInvariantError("invalid birthdate")
	}

	err = u.UserRepository.InsertCustomer(ctx, tx, userId, registerCustomer.Name, registerCustomer.Address, registerCustomer.PhoneNumber, registerCustomer.Gender, userBirthdate)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) RegisterSeller(ctx *gin.Context, registerSeller httpCommon.RegisterSeller) error {
	tx, err := u.DB.Begin()

	if err != nil {
		return errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	user, err := u.UserRepository.GetUserByEmail(ctx, tx, registerSeller.Email)

	if err != nil {
		return err
	}

	if user.UserID != "" {
		return errorCommon.NewInvariantError("email already registered")
	}

	hashedPassword, err := u.PasswordHashManager.HashPassword(registerSeller.Password)

	if err != nil {
		return errorCommon.NewInvariantError("failed to hash password")
	}

	userId, err := u.UserRepository.InsertUser(ctx, tx, registerSeller.Email, hashedPassword)

	if err != nil {
		return err
	}

	userBirthdate, err := time.Parse("2006-01-02", registerSeller.BirthDate)

	if err != nil {
		return errorCommon.NewInvariantError("invalid birthdate")
	}

	err = u.UserRepository.InsertSeller(ctx, tx, userId, registerSeller.Name, registerSeller.Address, registerSeller.PhoneNumber, registerSeller.Gender, registerSeller.IdentityNumber, registerSeller.BankName, registerSeller.DebitNumber, userBirthdate)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) GetCustomerByID(ctx *gin.Context, customerID string) (httpCommon.Customer, error) {
	tx, err := u.DB.Begin()

	if err != nil {
		return httpCommon.Customer{}, errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	customer, err := u.UserRepository.GetCustomerByID(ctx, tx, customerID)

	if err != nil {
		return httpCommon.Customer{}, err
	}

	user, err := u.UserRepository.GetUserByID(ctx, tx, customer.UserID)

	if err != nil {
		return httpCommon.Customer{}, err
	}

	return httpCommon.Customer{
		UserID:      customer.UserID,
		CustomerID:  customer.CustomerID,
		Email:       user.Email,
		Name:        customer.Name,
		Address:     customer.Address,
		PhoneNumber: customer.PhoneNumber,
		BirthDate:   customer.BirthDate.Format("2006-01-02"),
		Gender:      customer.Gender,
	}, nil

}

func (u *UserUsecaseImplementation) GetSellerByID(ctx *gin.Context, sellerID string) (httpCommon.Seller, error) {
	tx, err := u.DB.Begin()

	if err != nil {
		return httpCommon.Seller{}, errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	seller, err := u.UserRepository.GetSellerByID(ctx, tx, sellerID)

	if err != nil {
		return httpCommon.Seller{}, err
	}

	user, err := u.UserRepository.GetUserByID(ctx, tx, seller.UserID)

	if err != nil {
		return httpCommon.Seller{}, err
	}

	return httpCommon.Seller{
		UserID:         seller.UserID,
		SellerID:       seller.SellerID,
		Email:          user.Email,
		Name:           seller.Name,
		Address:        seller.Address,
		PhoneNumber:    seller.PhoneNumber,
		BirthDate:      seller.BirthDate.Format("2006-01-02"),
		Gender:         seller.Gender,
		IdentityNumber: seller.IdentityNumber,
		BankName:       seller.BankName,
		DebitNumber:    seller.DebitNumber,
	}, nil
}

func (u *UserUsecaseImplementation) ForgotPassword(ctx *gin.Context, forgotPassword httpCommon.ForgotPassword) error {
	return nil
}

func (u *UserUsecaseImplementation) ResetPassword(ctx *gin.Context, id, token string) error {
	return nil
}

func (u *UserUsecaseImplementation) ChangePassword(ctx *gin.Context, id string, ChangePassword httpCommon.ChangePassword) error {
	return nil
}

func (u *UserUsecaseImplementation) UpdateCustomerByID(ctx *gin.Context, customer httpCommon.Customer) error {
	return nil
}

func (u *UserUsecaseImplementation) UpdateSellerByID(ctx *gin.Context, seller httpCommon.Seller) error {
	return nil
}
