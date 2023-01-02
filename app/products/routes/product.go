package routes

import (
	"go-rest-api-mongo/app/products/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/product/:productId", controllers.FindOne)
	router.POST("/product", controllers.Create)
}
