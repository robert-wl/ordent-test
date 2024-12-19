package service

import (
	"fmt"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
)

type AuthService interface {
	LogIn(dto *dto.LogInRequest) (*string, error)
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
	_, err := s.repo.FindByEmailAndPassword(dto.Email, dto.Password)

	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	token := "token"
	return &token, nil
}
