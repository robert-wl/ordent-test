package dto

type CreateCommentRequest struct {
	ArticleID *string `json:"article_id"`
	ParentID  *string `json:"parent_id"`
	Title     string  `json:"title" binding:"required"`
	Body      string  `json:"body" binding:"required"`
}
