package repository

import (
	"context"
	"step/repository/users"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetUsername(ctx context.Context, userID int) (string, error)
}

type Repository struct {
	UsersRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UsersRepository: users.NewUsersRepository(db),
	}
}
