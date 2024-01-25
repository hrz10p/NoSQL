package main

import (
	"fmt"
	"main/pkg/cookies"
	"main/pkg/models"
	"main/pkg/services"
	"net/http"
	"text/template"
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
		file := "./ui/login.html"
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := handler.Service.EmployerService.GetByUsername(username)
		fmt.Println(user)
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

		cookies.SetCookie(w, user.ID.Hex(), t, "session")

		http.Redirect(w, r, "/employer", http.StatusSeeOther)

	}
}

func (handler *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookies.DeleteCookie(w, "session")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (handler *AuthHandler) RegisterAsEmployer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		file := "./ui/register.html"
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
