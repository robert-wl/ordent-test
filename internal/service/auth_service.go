package service

import (
	"fmt"
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
	"ordent-test/pkg/auth"
	"ordent-test/pkg/utils"
)

type AuthService interface {
	LogIn(dto *dto.LogInRequest) (*string, error)
	Register(dto *dto.RegisterRequest) error
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(r repository.UserRepository) AuthService {
	return &authService{
		repo: r,
	}
}

func (s *authService) LogIn(dto *dto.LogInRequest) (*string, error) {
	user, err := s.repo.FindByEmail(dto.Email)

	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if !utils.Compare(user.Password, dto.Password) {
		return nil, fmt.Errorf("password is incorrect")
	}

	token, err := auth.CreateJWT(user.SecureID)

	if err != nil {
		return nil, fmt.Errorf("failed to create token")
	}

	return token, nil
}

func (s *authService) Register(dto *dto.RegisterRequest) error {
	encryptPassword, err := utils.Encrypt(dto.Password)

	if err != nil {
		return fmt.Errorf("failed to encrypt password")
	}

	user := &model.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: encryptPassword,
	}

	err = s.repo.Create(user)

	if err != nil {
		return fmt.Errorf("failed to create user")
	}

	return nil
}
