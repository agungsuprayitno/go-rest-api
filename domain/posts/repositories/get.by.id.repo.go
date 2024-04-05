package repositories

import (
	"go-rest-postgres/domain/posts/models"
	errorhandlers "go-rest-postgres/error-handlers"
	"go-rest-postgres/initializers"

	"github.com/gin-gonic/gin"
)

func GetById(ctx *gin.Context) (postModel models.Post, err error) {
	postId := ctx.Param("postId")

	var post models.Post
	result := initializers.DB.Preload("UserData").First(&post, "id = ?", postId)
	if result.Error != nil {
		notfoundErr := errorhandlers.NotfoundError{}
		notfoundErr.SetError(ctx, "No post with that id exists")
		err = result.Error
	}
	return post, err
}
