package middleware

import (
	"fmt"
	"strings"

	"go-rest-postgres/domain/users/models"
	errorhandlers "go-rest-postgres/error-handlers"
	"go-rest-postgres/initializers"
	"go-rest-postgres/utils"

	"github.com/gin-gonic/gin"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")
		unauthorizedErr := errorhandlers.UnauthorizedError{}

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			unauthorizedErr.SetError(ctx, "You are not logged in")
			return
		}

		config, _ := initializers.LoadConfig()
		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			unauthorizedErr.SetError(ctx, err.Error())
			return
		}

		var user models.User
		result := initializers.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			unauthorizedErr.SetError(ctx, "the user belonging to this token no logger exists")
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
