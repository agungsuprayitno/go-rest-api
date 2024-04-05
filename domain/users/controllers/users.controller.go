package controllers

import (
	"go-rest-postgres/domain/users/responses"
	"go-rest-postgres/domain/users/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) GetByLoggedIn(ctx *gin.Context) {
	user := services.GetByLoggedIn(ctx)
	response := responses.UserResponse{}
	mappedResponse := response.MapResponse(user)
	ctx.JSON(http.StatusOK, gin.H{"data": mappedResponse})
}
