package services

import (
	"go-rest-api-mongo/app/products/models"
	"go-rest-api-mongo/app/products/repositories"
	"go-rest-api-mongo/app/products/requests"
)

func Create(request requests.Product) models.Product {
	return repositories.Create(request)
}
