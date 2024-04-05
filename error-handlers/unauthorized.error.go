package errorhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UnauthorizedError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (handler UnauthorizedError) SetError(ctx *gin.Context, message string) {
	errorHandler := UnauthorizedError{
		Status:  http.StatusUnauthorized,
		Message: message,
		Code:    "unauthorized",
	}

	ctx.AbortWithStatusJSON(errorHandler.Status, gin.H{"error": errorHandler})
	return
}
