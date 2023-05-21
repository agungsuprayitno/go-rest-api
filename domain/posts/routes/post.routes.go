package routes

import (
	"go-rest-postgres/domain/posts/controllers"
	"go-rest-postgres/middleware"

	"github.com/gin-gonic/gin"
)

type PostRouteController struct {
	postController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}

func (prc *PostRouteController) PostRoute(rg *gin.RouterGroup) {

	router := rg.Group("posts")
	router.Use(middleware.DeserializeUser())
	router.GET("/", prc.postController.GetPagination)
	router.GET("/all", prc.postController.GetAll)
	router.GET("/:postId", prc.postController.GetById)
	router.POST("/", prc.postController.Create)
	router.PUT("/:postId", prc.postController.Update)
	router.DELETE("/:postId", prc.postController.Delete)
	router.POST("/kafka-post", prc.postController.AsynchronousInsert)
}
