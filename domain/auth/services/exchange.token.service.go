package services

import (
	"go-rest-postgres/domain/auth/repositories"
	"go-rest-postgres/domain/auth/responses"

	"github.com/gin-gonic/gin"
)

func ExchangeToken(ctx *gin.Context) (tokenResp responses.TokenResponse, err error) {
	tokenResponse, err := repositories.ExchangeToken(ctx)

	return tokenResponse, err
}
