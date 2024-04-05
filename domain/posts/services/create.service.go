package services

import (
	"go-rest-postgres/domain/posts/models"
	"go-rest-postgres/domain/posts/repositories"
	"go-rest-postgres/domain/posts/requests"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) (postModel models.Post, err error) {

	var postRequest requests.CreatePostRequest
	ctx.BindJSON(&postRequest)

	var newPost models.Post
	newPost, err = repositories.Create(ctx, postRequest)

	return newPost, err
}
