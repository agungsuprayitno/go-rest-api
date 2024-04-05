package repositories

import (
	"go-rest-postgres/domain/posts/models"
	errorhandlers "go-rest-postgres/error-handlers"
	"go-rest-postgres/helpers"
	"go-rest-postgres/initializers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPagination(ctx *gin.Context) (postsPerPage []models.Post, pagination helpers.Pagination, err error) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	orderBy := ctx.DefaultQuery("order_by", "created_at")
	orderType := ctx.DefaultQuery("order_type", "desc")
	searchKeyword := ctx.DefaultQuery("search", "")

	var posts []models.Post
	var query = initializers.DB.Preload("UserData").Order(orderBy + " " + orderType).Find(&posts)
	if searchKeyword != "" {
		query.Where("title = ? ", searchKeyword)
	}
	total := query.RowsAffected
	pagination = pagination.SetPagination(page, perPage, total)

	offset := (page - 1) * perPage
	results := query.Limit(perPage).Offset(offset).Find(&posts)

	if results.Error != nil {
		badGatewayErr := errorhandlers.BadGatewayError{}
		badGatewayErr.SetError(ctx, results.Error.Error())
		err = results.Error
		return
	}

	return posts, pagination, err
}
