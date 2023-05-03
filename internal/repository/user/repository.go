package user

import (
	"context"
	"database/sql"
	"time"

	userDomain "github.com/aziemp66/byte-bargain/internal/domain/user"
)

type Repository interface {
	GetUserByID(ctx context.Context, tx *sql.Tx, userID string) (userDomain.User, error)
	GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (userDomain.User, error)
	GetCustomerByID(ctx context.Context, tx *sql.Tx, customerID string) (userDomain.Customer, error)
	GetSellerByID(ctx context.Context, tx *sql.Tx, sellerID string) (userDomain.Seller, error)
	GetCustomerByUserID(ctx context.Context, tx *sql.Tx, userID string) (userDomain.Customer, error)
	GetSellerByUserID(ctx context.Context, tx *sql.Tx, userID string) (userDomain.Seller, error)
	InsertUser(ctx context.Context, tx *sql.Tx, email, password string) (userID string, Error error)
	InsertCustomer(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender string, birthdate time.Time) error
	InsertSeller(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender, identityNumber, bankName, debitNumber string, birthdate time.Time) error
	UpdateCustomerByID(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender string, birthdate time.Time) error
	UpdateSellerByID(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender, identityNumber, bankName, debitNumber string, birthdate time.Time) error
	UpdateUserPasswordByID(ctx context.Context, tx *sql.Tx, userID string, password string) error
	UpdateUserVerifiedStatusByID(ctx context.Context, tx *sql.Tx, userID string, verified bool) error
}
