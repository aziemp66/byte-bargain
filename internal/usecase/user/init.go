package user

import (
	"database/sql"

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

func (u *UserUsecaseImplementation) Login() {

}
