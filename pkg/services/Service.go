package services

import "go.mongodb.org/mongo-driver/mongo"

type Service struct {
	EmployerService EmployerService
	VacancyService  VacancyService
}

func NewService(client *mongo.Database) *Service {
	return &Service{
		EmployerService: *NewEmployerService(client),
		VacancyService:  *NewVacancyService(client),
	}
}
