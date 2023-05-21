package userinfrastructure

import (
	"context"

	sqlc "github.com/Hirochon/infra-go-for-test/internal/infrastructure/sqlc/sqlcgenmodel"
)

func (u userRepository) LookUp(ctx context.Context, id string) (string, error) {
	querier := sqlc.New(u.mysqlClient)
	userID, err := querier.GetUser(ctx, id)
	if err != nil {
		return "", err
	}
	return userID, nil
}
