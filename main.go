package main

import (
	"go-rest-api-mongo/app/products/routes"
	"go-rest-api-mongo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()

	routes.ProductRoutes(router)
	router.Run("localhost: 3000")
}
