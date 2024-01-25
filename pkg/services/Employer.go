package services

import "go.mongodb.org/mongo-driver/mongo"

type EmployerService struct {
	client *mongo.Database
}

func NewEmployerService(client *mongo.Database) *EmployerService {
	return &EmployerService{client: client}
}

func (s *EmployerService) create() {

}

func (s *EmployerService) update() {

}

func (s *EmployerService) delete() {

}

func (s *EmployerService) get() {

}

func (s *EmployerService) getAll() {

}
