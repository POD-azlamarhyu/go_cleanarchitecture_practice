package user

import "context"

type IUserRepository interface {
	Save(ctx context.Context, u *User) error
}