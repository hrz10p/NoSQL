package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employer struct {
	ID          primitive.ObjectID `bson:"id"`
	Username    string             `bson:"username"`
	Password    string             `bson:"password"`
	CompanyName string             `bson:"company_name"`
	Vacancies   []Vacancy          `bson:"vacancies"`
}
