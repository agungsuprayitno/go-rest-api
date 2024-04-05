package controllers

import (
	"go-rest-postgres/domain/posts/responses"
	"go-rest-postgres/domain/posts/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (pc PostController) GetAll(ctx *gin.Context) {
	posts, err := services.GetAll(ctx)
	response := responses.PostResponse{}
	mappedResponses := response.MapResponses(posts)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": mappedResponses})
	}
}

func (pc PostController) GetPagination(ctx *gin.Context) {
	posts, pagination, err := services.GetPagination(ctx)
	response := responses.PostResponse{}
	mappedResponses, meta := response.MapPaginationResponses(posts, pagination)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": mappedResponses, "meta": meta})
	}
}

func (pc PostController) GetById(ctx *gin.Context) {
	post, err := services.GetById(ctx)
	response := responses.PostResponse{}
	mappedResponse := response.MapResponse(post)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": mappedResponse})
	}
}

func (pc PostController) Create(ctx *gin.Context) {
	newPost, err := services.Create(ctx)
	response := responses.PostResponse{}
	mappedResponse := response.MapResponse(newPost)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": mappedResponse})
	}
}

func (pc PostController) Update(ctx *gin.Context) {
	updatedPost, err := services.Update(ctx)

	response := responses.PostResponse{}
	mappedResponse := response.MapResponse(updatedPost)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": mappedResponse})
	}
}

func (pc PostController) Delete(ctx *gin.Context) {
	deletedPost, err := services.Delete(ctx)
	response := responses.PostResponse{}
	mappedResponse := response.MapResponse(deletedPost)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": mappedResponse})
	}
}
