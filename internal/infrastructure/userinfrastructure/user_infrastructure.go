package userinfrastructure

import (
	"database/sql"

	"github.com/Hirochon/infra-go-for-test/internal/domain/user"
)

type userRepository struct {
	mysqlClient *sql.DB
}

func NewUserRepository(mysqlClient *sql.DB) user.IUserRepository {
	return &userRepository{
		mysqlClient: mysqlClient,
	}
}
