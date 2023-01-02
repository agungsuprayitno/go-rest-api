package controllers

import (
	"go-rest-api-mongo/app/products/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindOne(c *gin.Context) {
	productIdParam := c.Param("productId")
	productId, _ := primitive.ObjectIDFromHex(productIdParam)
	product, err := services.FindOne(productId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": product})
}
