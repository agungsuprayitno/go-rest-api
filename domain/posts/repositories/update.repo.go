package repositories

import (
	"go-rest-postgres/domain/posts/models"
	"go-rest-postgres/domain/posts/requests"
	userModels "go-rest-postgres/domain/users/models"
	errorhandlers "go-rest-postgres/error-handlers"
	"go-rest-postgres/initializers"
	"time"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context, request requests.UpdatePostRequest) (postModel models.Post, err error) {
	postId := ctx.Param("postId")
	currentUser := ctx.MustGet("currentUser").(userModels.User)

	var updatedPost models.Post
	result := initializers.DB.First(&updatedPost, "id = ?", postId)
	if result.Error != nil {
		notfoundErr := errorhandlers.NotfoundError{}
		notfoundErr.SetError(ctx, "No post with that id exists")
		err = result.Error
		return
	}

	postToUpdate := models.Post{
		Title:     request.Title,
		Content:   request.Content,
		Image:     request.Image,
		User:      currentUser.ID,
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: time.Now(),
	}

	initializers.DB.Model(&updatedPost).Updates(postToUpdate)

	return postToUpdate, err
}
