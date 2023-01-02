package services

import (
	"go-rest-api-mongo/app/products/models"
	"go-rest-api-mongo/app/products/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindOne(productId primitive.ObjectID) (p models.Product, err error) {
	product, err := repositories.FindOne(productId)

	return product, err
}
