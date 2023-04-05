package user

import (
	"context"
	"database/sql"

	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"

	userRepository "github.com/aziemp66/byte-bargain/internal/repository/user"
)

type UserUsecaseImplementation struct {
	UserRepository userRepository.Repository
	DB             *sql.DB
	SessionManager *sessionCommon.SessionManager
}

func NewUserUsecaseImplementation(userRepository userRepository.Repository, db *sql.DB, sessionManager *sessionCommon.SessionManager) *UserUsecaseImplementation {
	return &UserUsecaseImplementation{
		UserRepository: userRepository,
		DB:             db,
		SessionManager: sessionManager,
	}
}

func (u *UserUsecaseImplementation) Login(ctx context.Context, login httpCommon.Login) error {
	return nil

}

func (u *UserUsecaseImplementation) RegisterCustomer(ctx context.Context, registerCustomer httpCommon.RegisterCustomer) error {
	return nil
}

func (u *UserUsecaseImplementation) RegisterSeller(ctx context.Context, registerSeller httpCommon.RegisterSeller) error {
	return nil
}

func (u *UserUsecaseImplementation) GetCustomerByID(ctx context.Context, customerID string) (httpCommon.Customer, error) {
	return httpCommon.Customer{}, nil
}

func (u *UserUsecaseImplementation) GetSellerByID(ctx context.Context, sellerID string) (httpCommon.Seller, error) {
	return httpCommon.Seller{}, nil
}

func (u *UserUsecaseImplementation) ForgotPassword(ctx context.Context, forgotPassword httpCommon.ForgotPassword) error {
	return nil
}

func (u *UserUsecaseImplementation) ResetPassword(ctx context.Context, id, token string) error {
	return nil
}

func (u *UserUsecaseImplementation) ChangePassword(ctx context.Context, id string, ChangePassword httpCommon.ChangePassword) error {
	return nil
}

func (u *UserUsecaseImplementation) UpdateCustomerByID(ctx context.Context, customer httpCommon.Customer) error {
	return nil
}

func (u *UserUsecaseImplementation) UpdateSellerByID(ctx context.Context, seller httpCommon.Seller) error {
	return nil
}
