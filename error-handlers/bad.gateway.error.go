package errorhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BadGatewayError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (handler BadGatewayError) SetError(ctx *gin.Context, message string) {
	errorHandler := BadGatewayError{
		Status:  http.StatusBadGateway,
		Message: message,
		Code:    "bad-gateway",
	}

	ctx.AbortWithStatusJSON(errorHandler.Status, gin.H{"error": errorHandler})
	return
}
