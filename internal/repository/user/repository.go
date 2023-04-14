package user

import (
	"database/sql"
	"time"

	userDomain "github.com/aziemp66/byte-bargain/internal/domain/user"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetUserByID(ctx *gin.Context, tx *sql.Tx, userID string) (userDomain.User, error)
	GetUserByEmail(ctx *gin.Context, tx *sql.Tx, email string) (userDomain.User, error)
	GetCustomerByID(ctx *gin.Context, tx *sql.Tx, customerID string) (userDomain.Customer, error)
	GetSellerByID(ctx *gin.Context, tx *sql.Tx, sellerID string) (userDomain.Seller, error)
	InsertUser(ctx *gin.Context, tx *sql.Tx, email, password string) (userID string, Error error)
	InsertCustomer(ctx *gin.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender string, birthdate time.Time) error
	InsertSeller(ctx *gin.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender, identityNumber, bankName, debitNumber string, birthdate time.Time) error
	UpdateCustomerByID(ctx *gin.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender string, birthdate time.Time) error
	UpdateSellerByID(ctx *gin.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender, identityNumber, bankName, debitNumber string, birthdate time.Time) error
	UpdateUserPasswordByID(ctx *gin.Context, tx *sql.Tx, userID string, password string) error
}
