package service

import (
	"fmt"
	"net/http"
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
	"ordent-test/pkg/auth"
	"ordent-test/pkg/utils"
	"strings"
)

type AuthService interface {
	LogIn(dto *dto.LogInRequest) (*string, error)
	Register(dto *dto.RegisterRequest) (*model.User, error)
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
		return nil, utils.NewAppError(
			fmt.Errorf("user with email %s not found", dto.Email),
			http.StatusNotFound,
			"user not found",
		)
	}

	if !utils.Compare(user.Password, dto.Password) {
		return nil, utils.NewAppError(
			fmt.Errorf("invalid password"),
			http.StatusUnauthorized,
			"invalid password",
		)
	}

	token, err := auth.CreateJWT(user)

	if err != nil {
		return nil, utils.NewAppError(
			fmt.Errorf("failed to create token"),
			http.StatusInternalServerError,
			"failed to create token",
		)
	}

	return token, nil
}

func (s *authService) Register(dto *dto.RegisterRequest) (*model.User, error) {
	encryptPassword, err := utils.Encrypt(dto.Password)

	if err != nil {
		return nil, utils.NewAppError(
			fmt.Errorf("failed to encrypt password"),
			http.StatusInternalServerError,
			"failed to encrypt password",
		)
	}

	user := &model.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: encryptPassword,
	}

	user, err = s.repo.Create(user)

	if err != nil {
		if strings.Contains(err.Error(), "idx_users_email") {
			return nil, utils.NewAppError(
				fmt.Errorf("email already exists"),
				http.StatusBadRequest,
				"email already exists",
			)
		}
		if strings.Contains(err.Error(), "idx_users_username") {
			return nil, utils.NewAppError(
				fmt.Errorf("username already exists"),
				http.StatusBadRequest,
				"username already exists",
			)
		}
		return nil, utils.NewAppError(
			err,
			http.StatusInternalServerError,
			"failed to create user",
		)
	}

	return user, nil
}
