package service

import (
	"context"
	"step/repository"
	"step/service/users"

	"gorm.io/gorm"
)

type UsersService interface {
	GetUsername(ctx context.Context, userID int) (string, error)
}

type Service struct {
	UsersService
}

func NewUsersService(db *gorm.DB, repo *repository.Repository) *Service {
	return &Service{
		UsersService: users.NewUsersService(repo),
	}
}
