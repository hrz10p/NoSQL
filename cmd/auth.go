package main

import (
	"main/pkg/cookies"
	"main/pkg/models"
	"main/pkg/services"
	"net/http"
	"time"
)

type AuthHandler struct {
	Service *services.Service
}

func NewAuthHandler(service *services.Service) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Render the login page
		// ...
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		userType := r.FormValue("userType")

		if userType == "employer" {

			user, err := handler.Service.EmployerService.GetByUsername(username)

			if err != nil {
				if err == models.ErrNoRecord {
					http.Error(w, "User not found", http.StatusBadRequest)
					return
				} else {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			if user.Password != password {
				http.Error(w, "Wrong password", http.StatusBadRequest)
				return
			}

			t := time.Now().Add(1 * time.Hour)

			cookies.SetCookie(w, user.ID.String(), t, "session")

			http.Redirect(w, r, "/applicant", http.StatusSeeOther)

		} else if userType == "applicant" {
			user, err := handler.Service.ApplicantService.GetByUsername(username)

			if err != nil {
				if err == models.ErrNoRecord {
					http.Error(w, "User not found", http.StatusBadRequest)
					return
				} else {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			if user.Password != password {
				http.Error(w, "Wrong password", http.StatusBadRequest)
				return
			}

			t := time.Now().Add(1 * time.Hour)

			cookies.SetCookie(w, user.ID.String(), t, "session")

			http.Redirect(w, r, "/employer", http.StatusSeeOther)
		}

	}
}

func (handler *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookies.DeleteCookie(w, "session")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (handler *AuthHandler) RegisterAsEmployer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Render the registration page
		// ...
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		obj := models.Employer{
			Username: username,
			Password: password,
		}

		err := handler.Service.EmployerService.Create(obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func (handler *AuthHandler) RegisterAsApplicant(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		// Render the registration page
		// ...
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		obj := models.Applicant{
			Username: username,
			Password: password,
		}

		err := handler.Service.ApplicantService.Create(obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
