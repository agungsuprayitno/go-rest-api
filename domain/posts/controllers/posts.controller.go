package controllers

import (
	"context"
	"fmt"
	"go-rest-postgres/domain/posts/responses"
	"go-rest-postgres/domain/posts/services"
	"go-rest-postgres/kafka/config"
	"go-rest-postgres/kafka/producer"
	"log"
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

func (pc PostController) AsynchronousInsert(ctx *gin.Context) {

	kafkaProducer, err := config.Configure("quickstart-events")
	if err != nil {
		log.Fatalln("error", err.Error())
		return
	}
	defer kafkaProducer.Close()

	err = producer.PushMessage(context.Background(), nil, []byte("test producer"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": map[string]interface{}{
				"message": fmt.Sprintf("error while push message into kafka: %s", err.Error()),
			},
		})

		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success push data into kafka",
		"data":    "test producer",
	})


	// deletedPost, err := services.Delete(ctx)
	// response := responses.PostResponse{}
	// mappedResponse := response.MapResponse(deletedPost)
	// if err == nil {
	// 	ctx.JSON(http.StatusOK, gin.H{"data": mappedResponse})
	// }
}
