package main

import (
	"main/pkg/models"
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

		user := getUserFromContext(r)

		err = tmpl.Execute(w, user)
		if err != nil {
			http.Error(w, "asd ", http.StatusInternalServerError)
			return
		}

	}

}

func (handler *EmployerHanlder) CreateVacancy(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")

		vac := models.Vacancy{
			Name:        name,
			Description: description,
		}

		err := handler.Service.VacancyService.Create(vac)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (handler *EmployerHanlder) GetVacancies(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		file := "./ui/index.html"
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		vacs, err := handler.Service.VacancyService.GetAll()
		if err != nil {
			http.Error(w, "Error vacs", 500)
			return
		}

		err = tmpl.Execute(w, vacs)
		if err != nil {
			http.Error(w, "asd ", http.StatusInternalServerError)
			return
		}

	}

}

func (handler *EmployerHanlder) DeleteVacancy(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		id := r.FormValue("id")

		objId, _ := primitive.ObjectIDFromHex(id)

		err := handler.Service.VacancyService.Delete(objId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employer", http.StatusSeeOther)
	}

}

func (handler *EmployerHanlder) UpdateVacancy(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")

		objId, _ := primitive.ObjectIDFromHex(id)

		vac := models.Vacancy{
			Name:        name,
			Description: description,
		}

		err := handler.Service.VacancyService.Update(objId, vac)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employer", http.StatusSeeOther)
	}

}
