package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Applicant struct {
	ID          primitive.ObjectID   `bson:"id"`
	Username    string               `bson:"username"`
	Password    string               `bson:"password"`
	Description string               `bson:"description"`
	Responses   []primitive.ObjectID `bson:"responses"`
}
