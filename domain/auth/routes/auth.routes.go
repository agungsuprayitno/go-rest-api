package routes

import (
	"go-rest-postgres/domain/auth/controllers"
	"go-rest-postgres/middleware"

	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh-token", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
	router.POST("/", rc.authController.GetAuthorizationCode)
	router.POST("/token", rc.authController.ExchangeToken)
}
