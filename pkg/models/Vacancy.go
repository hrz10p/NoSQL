package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vacancy struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}
