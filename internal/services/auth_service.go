package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
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
	user, err := s.UserRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *AuthService) UpdateProfile(user *models.User) (*models.User, error) {
	updatedUser, err := s.UserRepository.UpdateUser(*user)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}
