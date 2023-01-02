package controllers

import (
	"errors"
	"go-rest-api-mongo/app/errorhandlers"
	"go-rest-api-mongo/app/products/requests"
	"go-rest-api-mongo/app/products/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Create(c *gin.Context) {
	request := requests.Product{}

	if err := c.ShouldBindJSON(&request); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]errorhandlers.ErrorMessage, len(ve))
			for i, fe := range ve {
				out[i] = errorhandlers.ErrorMessage{fe.Field(), errorhandlers.GetValidationErrorMessage(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	newProduct := services.Create(request)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": newProduct})
}
