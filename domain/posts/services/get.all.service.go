package services

import (
	"go-rest-postgres/domain/posts/models"
	"go-rest-postgres/domain/posts/repositories"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) (postsModel []models.Post, err error) {
	posts, err := repositories.GetAll(ctx)
	return posts, err
}
