package errorhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BadRequestError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (handler BadRequestError) SetError(ctx *gin.Context, message string) {
	errorHandler := BadRequestError{
		Status:  http.StatusBadRequest,
		Message: message,
		Code:    "bad-request",
	}

	ctx.AbortWithStatusJSON(errorHandler.Status, gin.H{"error": errorHandler})
	return
}
