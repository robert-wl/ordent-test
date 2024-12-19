package service

import (
	"net/http"
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
	"ordent-test/pkg/pagination"
	"ordent-test/pkg/utils"
)

type UserService interface {
	GetUsers(dto *dto.GetUserRequest) ([]*model.User, error)
	GetAdmins(dto *dto.GetAdminRequest) ([]*model.User, error)
	GetUser(secureID string) (*model.User, error)
	ChangeRole(secureID string, role string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) GetUsers(dto *dto.GetUserRequest) ([]*model.User, error) {
	if dto.Search == nil {
		dto.Search = new(string)
	}

	if dto.Pagination == nil {
		dto.Pagination = new(pagination.Pagination)
	}

	return s.repo.Find(dto.Search, dto.Pagination)
}

func (s *userService) GetAdmins(dto *dto.GetAdminRequest) ([]*model.User, error) {
	if dto.Search == nil {
		dto.Search = new(string)
	}

	if dto.Pagination == nil {
		dto.Pagination = new(pagination.Pagination)
	}

	return s.repo.FindAdmins(dto.Search, dto.Pagination)
}

func (s *userService) GetUser(secureID string) (*model.User, error) {
	return s.repo.FindBySecureID(secureID)
}

func (s *userService) ChangeRole(secureID string, role string) (*model.User, error) {
	user, err := s.repo.FindBySecureID(secureID)

	if err != nil {
		return nil, utils.NewAppError(
			err,
			http.StatusNotFound,
			"Failed to find user",
		)
	}

	user.Role = role

	return s.repo.Update(user)
}
