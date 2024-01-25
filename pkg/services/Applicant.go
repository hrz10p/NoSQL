package services

import "go.mongodb.org/mongo-driver/mongo"

type ApplicantService struct {
	client *mongo.Database
}

func NewApplicantService(client *mongo.Database) *ApplicantService {
	return &ApplicantService{client: client}
}

func (s *ApplicantService) create() {

}

func (s *ApplicantService) update() {

}

func (s *ApplicantService) delete() {

}

func (s *ApplicantService) get() {

}

func (s *ApplicantService) getAll() {

}
