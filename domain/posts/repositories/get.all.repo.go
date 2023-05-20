package repositories

import (
	"go-rest-postgres/domain/posts/models"
	errorhandlers "go-rest-postgres/error-handlers"
	"go-rest-postgres/initializers"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) (postsModel []models.Post, err error) {
	var posts []models.Post
	orderBy := ctx.DefaultQuery("order_by", "created_at")
	orderType := ctx.DefaultQuery("order_type", "desc")
	searchKeyword := ctx.DefaultQuery("search", "")

	var query = initializers.DB.Preload("UserData").Order(orderBy + " " + orderType)
	if searchKeyword != "" {
		query.Where("title = ? ", searchKeyword)
	}
	results := query.Find(&posts)
	if results.Error != nil {
		badGatewayErr := errorhandlers.BadGatewayError{}
		badGatewayErr.SetError(ctx, results.Error.Error())
		err = results.Error
		return
	}
	return posts, err
}
