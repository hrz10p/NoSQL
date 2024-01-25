package main

import (
	"net/http"

	"main/pkg/services"
)

type ApplicantHandler struct {
	Service *services.Service
}

func NewApplicantHandler(service *services.Service) *ApplicantHandler {
	return &ApplicantHandler{
		Service: service,
	}
}

func (handler *ApplicantHandler) GetApplicant(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get applicant
}

func (handler *ApplicantHandler) CreateRespone(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get applicant
}
