package dto

type LogInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogInResponse struct {
	Token string `json:"token"`
}
