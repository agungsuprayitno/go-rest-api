package services

import (
	"go-rest-postgres/domain/posts/models"
	"go-rest-postgres/domain/posts/repositories"

	"github.com/gin-gonic/gin"
)

func GetById(ctx *gin.Context) (postModel models.Post, err error) {
	post, err := repositories.GetById(ctx)
	return post, err
}
