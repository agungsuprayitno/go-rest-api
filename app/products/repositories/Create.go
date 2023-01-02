package repositories

import (
	"context"
	"go-rest-api-mongo/app/products/collections"
	"go-rest-api-mongo/app/products/errors"
	"go-rest-api-mongo/app/products/models"
	"go-rest-api-mongo/app/products/requests"
	"go-rest-api-mongo/initializers"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(request requests.Product) models.Product {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = initializers.ConnectDB()
	var productCollection = collections.GetCollection(DB, "products")
	now := time.Now()

	newProduct := models.Product{
		Id:        primitive.NewObjectID(),
		Code:      request.Code,
		Name:      request.Name,
		CreatedAt: now.Format("2006-01-01"),
	}

	defer cancel()
	_, err := productCollection.InsertOne(ctx, newProduct)
	if err != nil {
		errors.NotFound(err)
	}
	return newProduct
}
