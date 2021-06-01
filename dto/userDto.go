package dto

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,notblank,max=200"`
	Password string `json:"password" binding:"required,notblank,max=200"`
}
