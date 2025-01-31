package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     *string            `bson:"name,omitempty"`
	Price    *int               `bson:"price,omitempty"`
	Category *string            `bson:"category,omitempty"`
}
