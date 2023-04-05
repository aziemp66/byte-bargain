package user

import (
	"context"
	"database/sql"
	userDomain "github.com/aziemp66/byte-bargain/internal/domain/user"
)

type UserRepositoryImplementation struct {
}

func NewUserRepositoryImplementation() *UserRepositoryImplementation {
	return &UserRepositoryImplementation{}
}

func (u UserRepositoryImplementation) GetUserByID(ctx context.Context, tx *sql.Tx, userID string) (userDomain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImplementation) GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (userDomain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImplementation) GetCustomerByID(ctx context.Context, tx *sql.Tx, customerID string) (userDomain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImplementation) GetSellerByID(ctx context.Context, tx *sql.Tx, sellerID string) (userDomain.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImplementation) InsertUser(ctx context.Context, tx *sql.Tx, email, password string) (userID string, Error error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImplementation) InsertCustomer(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImplementation) InsertSeller(ctx context.Context, tx *sql.Tx, userID, name, address, birthdate, phoneNumber, gender, identityNumber, bankName, debitNumber string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImplementation) UpdateCustomerByID(ctx context.Context, tx *sql.Tx, userID, name, address, phoneNumber, gender string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImplementation) UpdateSellerByID(ctx context.Context, tx *sql.Tx, userID, name, address, birthdate, phoneNumber, gender, identityNumber, bankName, debitNumber string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImplementation) UpdateUserPasswordByID(ctx context.Context, tx *sql.Tx, userID string, password string) error {
	//TODO implement me
	panic("implement me")
}
