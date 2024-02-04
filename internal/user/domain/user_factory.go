package domain

import (
	"Clean-Sweep-Solutions_/internal/utils"
	"time"
)

type UserFactory struct{}

func (f UserFactory) CreateAdmin(user *NewUser) *User {
	hashedPassword, _ := utils.HashPassword(user.Password)
	return &User{
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  hashedPassword,
		Role:      UserRole,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
	}
}

func (f UserFactory) CreateUser(user *NewUser) *User {
	hashedPassword, _ := utils.HashPassword(user.Password)
	return &User{
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  hashedPassword,
		Role:      AdminRole,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
	}
}
