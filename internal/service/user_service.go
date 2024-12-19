package service

import (
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
	"ordent-test/pkg/pagination"
)

type UserService interface {
	GetUsers(dto *dto.GetUserRequest) ([]*model.User, error)
	GetAdmins(dto *dto.GetAdminRequest) ([]*model.User, error)
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

	return s.repo.Find(dto.Search, dto.Pagination)
}
