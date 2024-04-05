package repositories

import (
	"go-rest-postgres/domain/auth/requests"
	"go-rest-postgres/domain/auth/responses"
	"go-rest-postgres/domain/users/models"
	errorhandlers "go-rest-postgres/error-handlers"
	"go-rest-postgres/initializers"
	"go-rest-postgres/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExchangeToken(ctx *gin.Context) (tokenResp responses.TokenResponse, signinErr error) {
	var payload *requests.SignInRequest

	badRequestErr := errorhandlers.BadRequestError{}
	forbiddenErr := errorhandlers.ForbiddenError{}

	if err1 := ctx.ShouldBindJSON(&payload); err1 != nil {
		badRequestErr.SetError(ctx, err1.Error())
		signinErr = err1
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		forbiddenErr.SetError(ctx, "Invalid Email or Password")
		signinErr = result.Error
		return
	}

	if err2 := utils.VerifyPassword(user.Password, payload.Password); err2 != nil {
		forbiddenErr.SetError(ctx, "Invalid Email or Password")
		signinErr = err2
		return
	}

	config, _ := initializers.LoadConfig()

	// Generate Tokens
	accessToken, err3 := utils.CreateToken(config.AccessTokenExpiresIn, user.Email, config.AccessTokenPrivateKey)
	if err3 != nil {
		badRequestErr.SetError(ctx, err3.Error())
		signinErr = err3
		return
	}

	refreshToken, err4 := utils.CreateToken(config.RefreshTokenExpiresIn, user.Email, config.RefreshTokenPrivateKey)
	if err4 != nil {
		badRequestErr.SetError(ctx, err4.Error())
		signinErr = err4
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	tokenResponse := responses.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		LoggedIn:     true,
		User:         user,
	}

	return tokenResponse, signinErr
}

func GenerateAuthorization(ctx *gin.Context) (authorizationResp responses.AuthorizationResponse, signinErr error) {
	var payload *requests.SignInRequest

	badRequestErr := errorhandlers.BadRequestError{}
	forbiddenErr := errorhandlers.ForbiddenError{}

	if err1 := ctx.ShouldBindJSON(&payload); err1 != nil {
		badRequestErr.SetError(ctx, err1.Error())
		signinErr = err1
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		forbiddenErr.SetError(ctx, "Invalid Email or Password")
		signinErr = result.Error
		return
	}

	if err2 := utils.VerifyPassword(user.Password, payload.Password); err2 != nil {
		forbiddenErr.SetError(ctx, "Invalid Email or Password")
		signinErr = err2
		return
	}

	config, _ := initializers.LoadConfig()

	// Generate Tokens
	accessToken, err3 := utils.CreateToken(config.AccessTokenExpiresIn, user.Email, config.AccessTokenPrivateKey)
	if err3 != nil {
		badRequestErr.SetError(ctx, err3.Error())
		signinErr = err3
		return
	}

	refreshToken, err4 := utils.CreateToken(config.RefreshTokenExpiresIn, user.Email, config.RefreshTokenPrivateKey)
	if err4 != nil {
		badRequestErr.SetError(ctx, err4.Error())
		signinErr = err4
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	authorizationResponse := responses.AuthorizationResponse{
		AuthorizationCode:  accessToken,
	}

	return authorizationResponse, signinErr
}
