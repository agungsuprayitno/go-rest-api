package services

import (
	"go-rest-postgres/domain/auth/repositories"
	"go-rest-postgres/domain/users/models"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) (userRsp models.User, err error) {
	user, err := repositories.SignUp(ctx)

	return user, err
}
