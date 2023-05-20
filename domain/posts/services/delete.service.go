package services

import (
	"go-rest-postgres/domain/posts/models"
	"go-rest-postgres/domain/posts/repositories"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) (postModel models.Post, err error) {
	deletedPost, err := repositories.Delete(ctx)

	return deletedPost, err
}
