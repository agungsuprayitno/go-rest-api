package main

import (
	"log"
	"net/http"

	authController "go-rest-postgres/domain/auth/controllers"
	authRoutes "go-rest-postgres/domain/auth/routes"
	postController "go-rest-postgres/domain/posts/controllers"
	postRoutes "go-rest-postgres/domain/posts/routes"
	userController "go-rest-postgres/domain/users/controllers"
	userRoutes "go-rest-postgres/domain/users/routes"
	"go-rest-postgres/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	AuthController      authController.AuthController
	AuthRouteController authRoutes.AuthRouteController

	UserController      userController.UserController
	UserRouteController userRoutes.UserRouteController

	PostController      postController.PostController
	PostRouteController postRoutes.PostRouteController
)

func init() {
	config, err := initializers.LoadConfig()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthRouteController = authRoutes.NewAuthRouteController(AuthController)
	UserRouteController = userRoutes.NewRouteUserController(UserController)
	PostRouteController = postRoutes.NewRoutePostController(PostController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
