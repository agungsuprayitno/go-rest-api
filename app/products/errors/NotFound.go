package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(err error) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	}
}
