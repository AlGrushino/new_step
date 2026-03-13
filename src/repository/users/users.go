package users

import (
	"context"

	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func (r *UsersRepository) GetUsername(ctx context.Context, userID int) (string, error) {
	var name string

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&name).Error

	if err != nil {
		return "", err
	}

	return name, nil
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}
