package main

import (
	"main/pkg/cookies"
	"main/pkg/services"
	"net/http"
	"text/template"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployerHanlder struct {
	Service *services.Service
}

func NewEmployerHandler(Service *services.Service) *EmployerHanlder {
	return &EmployerHanlder{
		Service: Service,
	}
}

func (handler *EmployerHanlder) GetEmployer(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		file := "./ui/templates/employer.html"
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			http.Error(w, "Error parsing templates", 500)
			return
		}

		cookie, err := cookies.GetCookie(r, "session")
		if err != nil {
			http.Error(w, "cookie issues", http.StatusInternalServerError)
			return
		}

		objId, _ := primitive.ObjectIDFromHex(cookie.Value)

		user, _ := handler.Service.EmployerService.Get(objId)

		err = tmpl.Execute(w, user)
		if err != nil {
			http.Error(w, "asd ", http.StatusInternalServerError)
			return
		}

	}

}

func (handler *EmployerHanlder) CreateVacancy(w http.ResponseWriter, r *http.Request) {

}

func (handler *EmployerHanlder) GetVacancies(w http.ResponseWriter, r *http.Request) {

}

func (handler *EmployerHanlder) DeleteVacancy(w http.ResponseWriter, r *http.Request) {

}

func (handler *EmployerHanlder) UpdateVacancy(w http.ResponseWriter, r *http.Request) {

}

func (handler *EmployerHanlder) GetApplicants(w http.ResponseWriter, r *http.Request) {

}
