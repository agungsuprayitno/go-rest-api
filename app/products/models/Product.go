package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Code      string             `json:"code"`
	Name      string             `json:"name"`
	CreatedAt string             `json:"createdAt"`
}
