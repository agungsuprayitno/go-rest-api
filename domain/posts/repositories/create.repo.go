package repositories

import (
	"go-rest-postgres/domain/posts/models"
	"go-rest-postgres/domain/posts/requests"
	userModels "go-rest-postgres/domain/users/models"
	errorhandlers "go-rest-postgres/error-handlers"
	"go-rest-postgres/initializers"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context, request requests.CreatePostRequest) (postModel models.Post, err error) {
	currentUser := ctx.MustGet("currentUser").(userModels.User)
	newPost := models.Post{
		Title:     request.Title,
		Content:   request.Content,
		Image:     request.Image,
		User:      currentUser.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := initializers.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			conflictErr := errorhandlers.ConflictError{}
			conflictErr.SetError(ctx, "Post with that title already exists")
			err = result.Error
			return
		}

		badGatewayErr := errorhandlers.BadGatewayError{}
		badGatewayErr.SetError(ctx, result.Error.Error())
		err = result.Error
		return
	}

	return newPost, err
}
