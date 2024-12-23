package dto

import "ordent-test/pkg/pagination"

type GetArticleCommentRequest struct {
	*pagination.Pagination
	Search *string `form:"search,omitempty" binding:"omitempty"`
}

type CreateCommentRequest struct {
	ArticleID *string `json:"article_id"`
	ParentID  *string `json:"parent_id"`
	Title     string  `json:"title" binding:"required"`
	Body      string  `json:"body" binding:"required"`
}

type UpdateCommentRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
