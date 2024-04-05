package errorhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForbiddenError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (handler ForbiddenError) SetError(ctx *gin.Context, message string) {
	errorHandler := ForbiddenError{
		Status:  http.StatusForbidden,
		Message: message,
		Code:    "forbidden",
	}

	ctx.AbortWithStatusJSON(errorHandler.Status, gin.H{"error": errorHandler})
	return
}
