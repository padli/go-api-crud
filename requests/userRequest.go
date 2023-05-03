package requests

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,emailExist"`
	Password string `json:"password" binding:"required"`
}
