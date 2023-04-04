package user

import (
	"database/sql"

	userRepository "github.com/aziemp66/byte-bargain/internal/repository/user"
)

type UserUsecaseImplementation struct {
	UserRepository userRepository.Repository
	DB             *sql.DB
}
