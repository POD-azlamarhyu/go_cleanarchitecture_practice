package user

import (
	"github.com/google/uuid"
	"time"
	"log/slog"
)

type User struct {
	userId uuid.UUID
	userName string
	email string
	password string
	icon string
	createdAt time.Time
	updatedAt time.Time
}

type Users []User

func NewUser(userName string, email string, password string) (*User, error) {
	slog.Info("Layer: Domain, Context: User, Creating new user",
		slog.String("userName",userName),
		slog.String("email",email),
	)

	return &User{
		userId:   uuid.New(),
		userName: userName,
		email:    email,
		password: password,
		icon:     string(""),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (u *User) GetUserId() uuid.UUID {
	return u.userId
}

func (u *User) GetUserName() string {
	return u.userName
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) getPassword() string {
	return u.password
}

func (u *User) GetIcon() string {
	return u.icon
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.updatedAt
}

