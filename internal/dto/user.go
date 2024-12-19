package dto

import "ordent-test/pkg/pagination"

type GetUserRequest struct {
	*pagination.Pagination
	Search *string `form:"search,omitempty" binding:"omitempty"`
}

type GetAdminRequest struct {
	*pagination.Pagination
	Search *string `form:"search,omitempty" binding:"omitempty"`
}
