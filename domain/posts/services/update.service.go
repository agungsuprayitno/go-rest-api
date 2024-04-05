package services

import (
	"go-rest-postgres/domain/posts/models"
	"go-rest-postgres/domain/posts/repositories"
	"go-rest-postgres/domain/posts/requests"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) (postModel models.Post, err error) {

	var postRequest requests.UpdatePostRequest
	ctx.BindJSON(&postRequest)

	post, err := repositories.Update(ctx, postRequest)
	return post, err
}
