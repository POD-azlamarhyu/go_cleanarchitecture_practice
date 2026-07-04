package user

import (
	"cleanarchitecture-practice/src/domain/user"
	"database/sql"
	"fmt"
	"log/slog"
	"context"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.IUserRepository{
	return &userRepository{db: db}
}

func (ur *userRepository) Save(ctx context.Context, user *user.User) error {
	slog.Info("Layer: Infrastructure, Context: User, Saving user to database",
		slog.String("userId",user.GetUserId().String()),
		slog.String("userName",user.GetUserName()),
		slog.String("email",user.GetEmail()),
	)

	// query := "INSERT INTO users (user_id, user_name, email, password, icon, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	fmt.Printf("実際には、SQLを実行する\n")
	return nil
}