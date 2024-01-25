package services

import "go.mongodb.org/mongo-driver/mongo"

type Service struct {
	ApplicantService ApplicantService
	EmployerService  EmployerService
	VacancyService   VacancyService
}

func NewService(client *mongo.Database) *Service {
	return &Service{
		ApplicantService: *NewApplicantService(client),
		EmployerService:  *NewEmployerService(client),
		VacancyService:   *NewVacancyService(client),
	}
}
