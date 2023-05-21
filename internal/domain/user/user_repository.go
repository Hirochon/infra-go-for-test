package user

import "context"

type IUserRepository interface {
	LookUp(ctx context.Context, id string) (string, error)
}
