package user

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"

	dbCommon "github.com/aziemp66/byte-bargain/common/db"
	errorCommon "github.com/aziemp66/byte-bargain/common/error"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	jwtCommon "github.com/aziemp66/byte-bargain/common/jwt"
	mailCommon "github.com/aziemp66/byte-bargain/common/mail"
	passwordCommon "github.com/aziemp66/byte-bargain/common/password"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"

	userRepository "github.com/aziemp66/byte-bargain/internal/repository/user"
)

type UserUsecaseImplementation struct {
	UserRepository      userRepository.Repository
	DB                  *sql.DB
	SessionManager      *sessionCommon.SessionManager
	PasswordHashManager *passwordCommon.PasswordHashManager
	JWTManager          *jwtCommon.JWTManager
	MailDialer          *gomail.Dialer
}

func NewUserUsecaseImplementation(
	userRepository userRepository.Repository,
	db *sql.DB,
	sessionManager *sessionCommon.SessionManager,
	passwordManager *passwordCommon.PasswordHashManager,
	jwtManager *jwtCommon.JWTManager,
	mailDialer *gomail.Dialer,
) *UserUsecaseImplementation {
	return &UserUsecaseImplementation{
		UserRepository:      userRepository,
		DB:                  db,
		SessionManager:      sessionManager,
		PasswordHashManager: passwordManager,
		JWTManager:          jwtManager,
		MailDialer:          mailDialer,
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

	customer, err := u.UserRepository.GetCustomerByUserID(ctx, tx, user.UserID)

	if err != nil {
		return err
	}

	if customer.CustomerID != "" {
		err = u.SessionManager.SetSessionValue(ctx, "customer_id", customer.CustomerID)

		if err != nil {
			return errorCommon.NewInvariantError("failed to set session value")
		}
	}

	seller, err := u.UserRepository.GetSellerByUserID(ctx, tx, user.UserID)

	if err != nil {
		return err
	}

	if seller.SellerID != "" {
		err = u.SessionManager.SetSessionValue(ctx, "seller_id", seller.SellerID)

		if err != nil {
			return errorCommon.NewInvariantError("failed to set session value")
		}
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

	err = u.SendActivationEmail(ctx, registerCustomer.Email)

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

	err = u.SendActivationEmail(ctx, registerSeller.Email)

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
	tx, err := u.DB.Begin()

	if err != nil {
		return errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	user, err := u.UserRepository.GetUserByEmail(ctx, tx, forgotPassword.Email)

	if err != nil {
		return err
	}

	if user.UserID == "" {
		return errorCommon.NewInvariantError("email not registered")
	}

	token, err := u.JWTManager.GenerateUserToken(user.UserID, 3*24*time.Hour)

	if err != nil {
		return errorCommon.NewInvariantError("failed to generate token")
	}

	mailPasswordReset := mailCommon.PasswordReset{
		Email: user.Email,
		Token: token,
	}

	mailTemplate, err := mailCommon.RenderPasswordResetTemplate(mailPasswordReset, ctx.GetString("web_url"))

	if err != nil {
		return err
	}

	message := mailCommon.NewMessage(ctx.GetString("web_email"), user.Email, "Reset Password", mailTemplate)

	err = u.MailDialer.DialAndSend(message)

	if err != nil {
		return errorCommon.NewInvariantError("failed to send email")
	}

	return nil
}

func (u *UserUsecaseImplementation) ResetPassword(ctx *gin.Context, resetPassword httpCommon.ResetPassword) error {
	tx, err := u.DB.Begin()

	if err != nil {
		return errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	claims, err := u.JWTManager.VerifyUserToken(resetPassword.Token)

	if err != nil {
		return errorCommon.NewInvariantError("invalid token")
	}

	user, err := u.UserRepository.GetUserByID(ctx, tx, claims.ID)

	if err != nil {
		return err
	}

	if user.UserID == "" {
		return errorCommon.NewInvariantError("user not found")
	}

	newPassword, err := u.PasswordHashManager.HashPassword(resetPassword.Password)

	if err != nil {
		return errorCommon.NewInvariantError("failed to hash password")
	}

	err = u.UserRepository.UpdateUserPasswordByID(ctx, tx, claims.ID, newPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) ChangePassword(ctx *gin.Context, ChangePassword httpCommon.ChangePassword) error {
	tx, err := u.DB.Begin()

	if err != nil {
		return errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	id, ok := u.SessionManager.GetSessionValue(ctx, "user_id").(string)

	if !ok {
		return errorCommon.NewInvariantError("invalid session")
	}

	user, err := u.UserRepository.GetUserByID(ctx, tx, id)

	if err != nil {
		return err
	}

	if user.UserID == "" {
		return errorCommon.NewInvariantError("user not found")
	}

	err = u.PasswordHashManager.CheckPasswordHash(ChangePassword.OldPassword, user.Password)

	if err != nil {
		return errorCommon.NewInvariantError("invalid old password")
	}

	newPassword, err := u.PasswordHashManager.HashPassword(ChangePassword.NewPassword)

	if err != nil {
		return errorCommon.NewInvariantError("failed to hash new password")
	}

	err = u.UserRepository.UpdateUserPasswordByID(ctx, tx, id, newPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) UpdateCustomerByID(ctx *gin.Context, customer httpCommon.UpdateCustomer) error {
	tx, err := u.DB.Begin()

	if err != nil {
		return errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	userId, ok := u.SessionManager.GetSessionValue(ctx, "user_id").(string)

	if !ok {
		return errorCommon.NewInvariantError("failed to get user id")
	}

	customerBirthdate, err := time.Parse("2006-01-02", customer.BirthDate)

	if err != nil {
		return errorCommon.NewInvariantError("invalid birthdate")
	}

	err = u.UserRepository.UpdateCustomerByID(ctx, tx, userId, customer.Name, customer.Address, customer.PhoneNumber, customer.Gender, customerBirthdate)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) UpdateSellerByID(ctx *gin.Context, seller httpCommon.UpdateSeller) error {
	tx, err := u.DB.Begin()

	if err != nil {
		return errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	userId, ok := u.SessionManager.GetSessionValue(ctx, "user_id").(string)

	if !ok {
		return errorCommon.NewInvariantError("failed to get user id")
	}

	sellerBirthdate, err := time.Parse("2006-01-02", seller.BirthDate)

	if err != nil {
		return errorCommon.NewInvariantError("invalid birthdate")
	}

	err = u.UserRepository.UpdateSellerByID(ctx, tx, userId, seller.Name, seller.Address, seller.PhoneNumber, seller.Gender, seller.IdentityNumber, seller.BankName, seller.DebitNumber, sellerBirthdate)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) ActivateAccount(ctx *gin.Context, token string) error {
	tx, err := u.DB.Begin()

	if err != nil {
		return errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	claims, err := u.JWTManager.VerifyUserToken(token)

	if err != nil {
		return errorCommon.NewInvariantError("invalid token")
	}

	user, err := u.UserRepository.GetUserByID(ctx, tx, claims.ID)

	if err != nil {
		return err
	}

	if user.UserID == "" {
		return errorCommon.NewInvariantError("user not found")
	}

	err = u.UserRepository.UpdateUserVerifiedStatusByID(ctx, tx, user.UserID, true)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) SendActivationEmail(ctx *gin.Context, email string) error {
	tx, err := u.DB.Begin()

	if err != nil {
		return errorCommon.NewInvariantError("failed to begin transaction")
	}

	defer dbCommon.CommitOrRollback(tx)

	user, err := u.UserRepository.GetUserByEmail(ctx, tx, email)

	if err != nil {
		return err
	}

	if user.UserID == "" {
		return errorCommon.NewInvariantError("email not registered")
	}

	token, err := u.JWTManager.GenerateUserToken(user.UserID, 3*24*time.Hour)

	if err != nil {
		return errorCommon.NewInvariantError("failed to generate token")
	}

	mailActivation := mailCommon.EmailVerification{
		Token: token,
	}

	mailTemplate, err := mailCommon.RenderEmailVerificationTemplate(mailActivation, ctx.GetString("web_url"))

	if err != nil {
		return err
	}

	message := mailCommon.NewMessage(ctx.GetString("web_email"), user.Email, "Activate Account", mailTemplate)

	err = u.MailDialer.DialAndSend(message)

	if err != nil {
		return errorCommon.NewInvariantError("failed to send email")
	}

	return nil
}
