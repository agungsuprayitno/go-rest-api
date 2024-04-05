package services

import (
	"go-rest-postgres/domain/users/models"

	"github.com/gin-gonic/gin"
)

func GetByLoggedIn(ctx *gin.Context) models.User {
	currentUser := ctx.MustGet("currentUser").(models.User)

	return currentUser

}
