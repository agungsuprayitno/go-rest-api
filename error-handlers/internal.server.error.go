package errorhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InternalServerError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (handler InternalServerError) SetError(ctx *gin.Context, message string) {
	errorHandler := InternalServerError{
		Status:  http.StatusInternalServerError,
		Message: message,
		Code:    "internal-server",
	}

	ctx.AbortWithStatusJSON(errorHandler.Status, gin.H{"error": errorHandler})
	return
}
