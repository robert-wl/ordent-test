package service

import (
	"fmt"
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
	"ordent-test/pkg/auth"
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
	user, err := s.repo.FindByEmailAndPassword(dto.Email, dto.Password)

	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	token, err := auth.CreateJWT(user.SecureID)

	if err != nil {
		return nil, fmt.Errorf("failed to create token")
	}

	return token, nil
}

func (s *authService) Register(dto *dto.RegisterRequest) error {
	user := &model.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}

	err := s.repo.Create(user)

	if err != nil {
		return fmt.Errorf("failed to create user")
	}

	return nil
}
