package repository

import (
	"errors"
	"living-library/internal/database"
	"living-library/internal/models"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(user models.User) (models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	if err := database.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user models.User) (models.User, error) {
	if err := database.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
