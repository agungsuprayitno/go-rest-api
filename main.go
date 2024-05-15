package main

import (
	// "fmt"
	"log"
	"net/http"

	authController "go-rest-postgres/domain/auth/controllers"
	//	Uncomment this model if you wanna run a migration
	// authModel "go-rest-postgres/domain/auth/models"
	authRoutes "go-rest-postgres/domain/auth/routes"
	postController "go-rest-postgres/domain/posts/controllers"

	//	Uncomment this model if you wanna run a migration
	// postModel "go-rest-postgres/domain/posts/models"
	postRoutes "go-rest-postgres/domain/posts/routes"
	userController "go-rest-postgres/domain/users/controllers"

	//	Uncomment this model if you wanna run a migration
	// userModel "go-rest-postgres/domain/users/models"
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

	//	Uncomment this section if you wanna run a migration
	// initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	// initializers.DB.AutoMigrate(&userModel.User{}, &postModel.Post{})
	// initializers.DB.AutoMigrate(&authModel.Authorization{}, &authModel.Merchant{}, &authModel.MerchantHistory{})
	// fmt.Println("👍 Migration complete")

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
