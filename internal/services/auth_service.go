package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"living-library/internal/database"
	"living-library/internal/models"
	"living-library/internal/repository"
)

type AuthService struct {
	UserRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepo}
}

func (s *AuthService) Register(user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	createdUser, err := s.UserRepository.CreateUser(*user)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	var u models.User
	if err := database.DB.Where("username = ?", username).First(&u).Error; err != nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("login or password incorrect")
	}
	return &u, nil
}

func (s *AuthService) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *AuthService) UpdateProfile(user *models.User) (*models.User, error) {
	updatedUser, err := s.UserRepository.UpdateUser(*user)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}
