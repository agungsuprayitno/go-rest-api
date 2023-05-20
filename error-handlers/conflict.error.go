package errorhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConflictError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (handler ConflictError) SetError(ctx *gin.Context, message string) {
	errorHandler := ConflictError{
		Status:  http.StatusConflict,
		Message: message,
		Code:    "conflict",
	}

	ctx.AbortWithStatusJSON(errorHandler.Status, gin.H{"error": errorHandler})
	return
}
