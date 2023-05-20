package services

import (
	"fmt"
	"go-rest-postgres/domain/auth/responses"
	"go-rest-postgres/domain/users/models"
	errorhandlers "go-rest-postgres/error-handlers"
	"go-rest-postgres/initializers"
	"go-rest-postgres/utils"

	"github.com/gin-gonic/gin"
)

// Refresh Access Token
func RefreshAccessToken(ctx *gin.Context) (token responses.TokenResponse, err error) {
	message := "could not refresh access token"

	refreshToken, err := ctx.Cookie("refresh_token")
	forbiddenErr := errorhandlers.ForbiddenError{}

	if err != nil {
		forbiddenErr.SetError(ctx, message)
		return
	}

	config, _ := initializers.LoadConfig()

	sub, err := utils.ValidateToken(refreshToken, config.RefreshTokenPublicKey)
	if err != nil {
		forbiddenErr.SetError(ctx, err.Error())
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "id = ?", fmt.Sprint(sub))
	if result.Error != nil {
		forbiddenErr.SetError(ctx, "the user belonging to this token no logger exists")
		return
	}

	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		forbiddenErr.SetError(ctx, err.Error())
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	tokenResponse := responses.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		LoggedIn:     true,
		User:         user,
	}

	return tokenResponse, nil

}
