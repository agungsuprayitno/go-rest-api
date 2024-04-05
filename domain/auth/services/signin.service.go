package services

import (
	"go-rest-postgres/domain/auth/repositories"
	"go-rest-postgres/domain/auth/responses"

	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) (tokenResp responses.TokenResponse, err error) {
	tokenResponse, err := repositories.SignIn(ctx)

	return tokenResponse, err
}
