package main

import (
	"fmt"
	"main/pkg"
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

		file := "./ui/employer.html"
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			http.Error(w, "Error parsing templates", 500)
			return
		}

		user := getUserFromContext(r)
		data := userConvert(user)
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println(err)
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

		newvac, err := handler.Service.VacancyService.Create(vac)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user := getUserFromContext(r)

		user.Vacancies = append(user.Vacancies, newvac)

		err = handler.Service.EmployerService.Update(user.ID, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employer", http.StatusSeeOther)
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

		user := getUserFromContext(r)

		index := -1
		for i, vacancy := range user.Vacancies {
			if vacancy.ID == objId {
				index = i
				break
			}
		}
		if index != -1 {
			user.Vacancies = append(user.Vacancies[:index], user.Vacancies[index+1:]...)
		}

		err := handler.Service.VacancyService.Delete(objId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = handler.Service.EmployerService.Update(user.ID, user)
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
		user := getUserFromContext(r)
		index := -1
		for i, vacancy := range user.Vacancies {
			if vacancy.ID == objId {
				index = i
				break
			}
		}
		user.Vacancies[index] = vac
		err = handler.Service.EmployerService.Update(user.ID, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employer", http.StatusSeeOther)
	}

}

// func (handler *EmployerHanlder) UpdateProfile(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodPost {

// 		id := r.FormValue("id")
// 		name := r.FormValue("name")
// 		description := r.FormValue("description")
// 		compName := r.FormValue("compName")
// 		objId, _ := primitive.ObjectIDFromHex(id)

// 		vac := models.Vacancy{
// 			Name:        name,
// 			Description: description,
// 		}

// 		err := handler.Service.VacancyService.Update(objId, vac)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		http.Redirect(w, r, "/employer", http.StatusSeeOther)
// 	}

// }

func userConvert(emp models.Employer) pkg.EmployerView {
	user := pkg.EmployerView{
		ID:          emp.ID.Hex(),
		Username:    emp.Username,
		Password:    emp.Password,
		CompanyName: emp.CompanyName,
	}
	var vacs []pkg.VacancyView

	for _, vac := range emp.Vacancies {
		vacs = append(vacs, pkg.VacancyView{
			ID:          vac.ID.Hex(),
			Name:        vac.Name,
			Description: vac.Description,
		})
	}

	user.Vacancies = vacs

	return user
}
