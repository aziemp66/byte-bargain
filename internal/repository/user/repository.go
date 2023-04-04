package user

import (
	"context"
	"database/sql"

	userDomain "github.com/aziemp66/byte-bargain/internal/domain/user"
)

type Repository interface {
	GetUserByID(ctx context.Context, tx *sql.Tx, userID string) (userDomain.User, error)
	GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (userDomain.User, error)
	GetCustomerByID(ctx context.Context, tx *sql.Tx, customerID string) (userDomain.Customer, error)
	GetSellerByID(ctx context.Context, tx *sql.Tx, sellerID string) (userDomain.Seller, error)
	InsertUser(ctx context.Context, tx *sql.Tx, email, password string) (userID string, Error error)
	InsertCustomer(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender string) error
	InsertSeller(ctx context.Context, tx *sql.Tx, userID, name, address, birthdate, phoneNumber, gender, identityNumber, bankName, debitNumber string) error
	UpdateCustomerByID(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender string) error
	UpdateSellerByID(ctx context.Context, tx *sql.Tx, userID, name, address, birthdate, phoneNumber, gender, identityNumber, bankName, debitNumber string) error
	UpdateUserPasswordByID(ctx context.Context, tx *sql.Tx, userID string, password string) error
}
