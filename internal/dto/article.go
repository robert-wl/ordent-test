package dto

import "ordent-test/pkg/pagination"

type GetArticleRequest struct {
	*pagination.Pagination
	Search *string `form:"search,omitempty" binding:"omitempty"`
}

type CreateArticleRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type UpdateArticleRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
