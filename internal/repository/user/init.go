package user

import (
	"context"
	"database/sql"
	"time"

	errorCommon "github.com/aziemp66/byte-bargain/common/error"
	userDomain "github.com/aziemp66/byte-bargain/internal/domain/user"
	"github.com/google/uuid"
)

type UserRepositoryImplementation struct {
}

func NewUserRepositoryImplementation() *UserRepositoryImplementation {
	return &UserRepositoryImplementation{}
}

func (u UserRepositoryImplementation) GetUserByID(ctx context.Context, tx *sql.Tx, userID string) (userDomain.User, error) {
	var user userDomain.User

	query := `SELECT user_id, email, password FROM user WHERE user_id = ?`

	err := tx.QueryRowContext(ctx, query, userID).Scan(&user.UserID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, errorCommon.NewInvariantError("user not found")
	}

	return user, nil
}

func (u UserRepositoryImplementation) GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (userDomain.User, error) {
	var user userDomain.User

	query := `SELECT user_id, email, password FROM user WHERE email = ?`

	err := tx.QueryRowContext(ctx, query, email).Scan(&user.UserID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, errorCommon.NewInvariantError("user not found")
	}

	return user, nil
}

func (u UserRepositoryImplementation) GetCustomerByID(ctx context.Context, tx *sql.Tx, customerID string) (userDomain.Customer, error) {
	var customer userDomain.Customer

	query := `SELECT customer_id, user_id, name, address, date_of_birth, phone_number, gender FROM customer WHERE customer_id = ?`

	err := tx.QueryRowContext(ctx, query, customerID).Scan(&customer.CustomerID, &customer.UserID, &customer.Name, &customer.Address, &customer.BirthDate, &customer.PhoneNumber, &customer.Gender)

	if err != nil {
		return customer, errorCommon.NewInvariantError("customer not found")
	}

	return customer, nil
}

func (u UserRepositoryImplementation) GetCustomerByUserID(ctx context.Context, tx *sql.Tx, userID string) (userDomain.Customer, error) {
	var customer userDomain.Customer

	query := `SELECT customer_id, user_id, name, address, date_of_birth, phone_number, gender FROM customer WHERE user_id = ?`

	err := tx.QueryRowContext(ctx, query, userID).Scan(&customer.CustomerID, &customer.UserID, &customer.Name, &customer.Address, &customer.BirthDate, &customer.PhoneNumber, &customer.Gender)

	if err != nil {
		return customer, errorCommon.NewInvariantError("customer not found")
	}

	return customer, nil
}

func (u UserRepositoryImplementation) GetSellerByID(ctx context.Context, tx *sql.Tx, sellerID string) (userDomain.Seller, error) {
	var seller userDomain.Seller

	query := `Select seller_id, user_id, name, address, date_of_birth, phone_number, gender, identity_number, bank_name, debit_number FROM seller WHERE seller_id = ?`

	err := tx.QueryRowContext(ctx, query, sellerID).Scan(&seller.SellerID, &seller.UserID, &seller.Name, &seller.Address, &seller.BirthDate, &seller.PhoneNumber, &seller.Gender, &seller.IdentityNumber, &seller.BankName, &seller.DebitNumber)

	if err != nil {
		return seller, errorCommon.NewInvariantError("seller not found")
	}

	return seller, nil
}

func (u UserRepositoryImplementation) GetSellerByUserID(ctx context.Context, tx *sql.Tx, userID string) (userDomain.Seller, error) {
	var seller userDomain.Seller

	query := `Select seller_id, user_id, name, address, date_of_birth, phone_number, gender, identity_number, bank_name, debit_number FROM seller WHERE user_id = ?`

	err := tx.QueryRowContext(ctx, query, userID).Scan(&seller.SellerID, &seller.UserID, &seller.Name, &seller.Address, &seller.BirthDate, &seller.PhoneNumber, &seller.Gender, &seller.IdentityNumber, &seller.BankName, &seller.DebitNumber)

	if err != nil {
		return seller, errorCommon.NewInvariantError("seller not found")
	}

	return seller, nil
}

func (u UserRepositoryImplementation) InsertUser(ctx context.Context, tx *sql.Tx, email, password string) (userID string, Error error) {
	query := `INSERT INTO user (user_id, email, password) VALUES (?, ?, ?)`

	userID = uuid.NewString()

	_, err := tx.ExecContext(ctx, query, userID, email, password)

	if err != nil {
		return "", errorCommon.NewInvariantError(err.Error())
	}

	return userID, nil
}

func (u UserRepositoryImplementation) InsertCustomer(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender string, birthdate time.Time) error {
	query := `INSERT INTO customer (customer_id, user_id, name, address, date_of_birth, phone_number, gender) VALUES (?, ?, ?, ?, ?, ?, ?)`

	customerID := uuid.NewString()

	_, err := tx.ExecContext(ctx, query, customerID, userID, name, address, birthdate, phoneNumber, gender)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (u UserRepositoryImplementation) InsertSeller(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender, identityNumber, bankName, debitNumber string, birthdate time.Time) error {
	query := `INSERT INTO seller (seller_id, user_id, name, address, date_of_birth, phone_number, gender, identity_number, bank_name, debit_number) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	sellerID := uuid.NewString()

	_, err := tx.ExecContext(ctx, query, sellerID, userID, name, address, birthdate, phoneNumber, gender, identityNumber, bankName, debitNumber)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (u UserRepositoryImplementation) UpdateCustomerByID(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender string, birthdate time.Time) error {
	query := `UPDATE customer set name = ?, address = ?, phone_number = ?, gender = ?, date_of_birth = ? WHERE customer_id = ?`

	_, err := tx.ExecContext(ctx, query, name, address, phoneNumber, gender, birthdate, userID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (u UserRepositoryImplementation) UpdateSellerByID(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender, identityNumber, bankName, debitNumber string, birthdate time.Time) error {
	query := `UPDATE seller set name = ?, address = ?, phone_number = ?, gender = ?, date_of_birth = ?, identity_number = ?, bank_name = ?, debit_number = ? WHERE customer_id = ?`

	_, err := tx.ExecContext(ctx, query, name, address, phoneNumber, gender, birthdate, identityNumber, bankName, debitNumber, userID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (u UserRepositoryImplementation) UpdateUserPasswordByID(ctx context.Context, tx *sql.Tx, userID string, password string) error {
	query := `UPDATE user set password = ? WHERE user_id = ?`

	_, err := tx.ExecContext(ctx, query, password, userID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (u UserRepositoryImplementation) UpdateUserVerifiedStatusByID(ctx context.Context, tx *sql.Tx, userID string, status bool) error {
	query := `UPDATE user set is_verified = ? WHERE user_id = ?`

	_, err := tx.ExecContext(ctx, query, status, userID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}
