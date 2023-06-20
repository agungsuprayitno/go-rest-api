package requests

import "mime/multipart"

type SignUpRequest struct {
	Name            string                `form:"name" binding:"required"`
	Email           string                `form:"email" binding:"required"`
	Password        string                `form:"password" binding:"required,min=8"`
	PasswordConfirm string                `form:"password_confirm" binding:"required"`
	Photo           *multipart.FileHeader `form:"photo" binding:"required"`
}
