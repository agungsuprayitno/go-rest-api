package repositories

import (
	"context"
	"fmt"
	"go-rest-api-mongo/app/products/collections"
	"go-rest-api-mongo/app/products/models"
	"go-rest-api-mongo/initializers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindOne(productId primitive.ObjectID) (p models.Product, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = initializers.ConnectDB()
	var productCollection = collections.GetCollection(DB, "products")

	var product models.Product

	defer cancel()
	resultErr := productCollection.FindOne(ctx, bson.M{"_id": productId}).Decode(&product)

	fmt.Println(resultErr)

	return product, resultErr
}
