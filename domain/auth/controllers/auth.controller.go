package controllers

import (
	"go-rest-postgres/domain/auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	user, err := services.SignUp(ctx)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func (ac *AuthController) SignInUser(ctx *gin.Context) {
	tokenResponse, err := services.SignIn(ctx)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": tokenResponse})
	}
}

func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {
	tokenResponse, err := services.RefreshAccessToken(ctx)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": tokenResponse})
	}
}

func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	services.LogoutUser(ctx)
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ac *AuthController) GetAuthorizationCode(ctx *gin.Context) {
	authorizationResponse, err := services.GenerateAuthorization(ctx)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": authorizationResponse})
	}
}

func (ac *AuthController) ExchangeToken(ctx *gin.Context) {
	tokenResponse, err := services.ExchangeToken(ctx)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": tokenResponse})
	}
}

