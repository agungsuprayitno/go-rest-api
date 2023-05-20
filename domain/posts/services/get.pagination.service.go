package services

import (
	"go-rest-postgres/domain/posts/models"
	"go-rest-postgres/domain/posts/repositories"
	"go-rest-postgres/helpers"

	"github.com/gin-gonic/gin"
)

func GetPagination(ctx *gin.Context) (postsModel []models.Post, paginator helpers.Pagination, err error) {
	posts, pagination, err := repositories.GetPagination(ctx)

	return posts, pagination, err
}
