package users

import (
	"context"
	"step/repository"
)

type UsersService struct {
	repository *repository.Repository
}

func NewUsersService(repository *repository.Repository) *UsersService {
	return &UsersService{
		repository: repository,
	}
}

func (s *UsersService) GetUsername(ctx context.Context, userID int) (string, error) {
	name, err := s.repository.GetUsername(ctx, userID)
	if err != nil {
		return "", err
	}

	return name, nil
}
