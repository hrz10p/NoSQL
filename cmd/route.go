package main

import (
	"fmt"
	"net/http"
)

// InitializeRoutes sets up the application routes
func (app *Application) InitializeRoutes() {

	users, err := app.Service.EmployerService.GetAll()

	if err != nil {
		app.Logger.Error(err.Error())
	}

	fmt.Println(users)

	emp := NewEmployerHandler(app.Service)
	auth := NewAuthHandler(app.Service)
	m := NewMiddle(app.Service)
	app.Router.HandleFunc("/", emp.GetVacancies)

	app.Router.HandleFunc("/login", auth.Login)
	app.Router.HandleFunc("/logout", auth.Logout)
	app.Router.HandleFunc("/register", auth.RegisterAsEmployer)

	app.Router.Handle("/employer", m.Authenticate(m.RequireAuthentication(http.HandlerFunc(emp.GetEmployer))))
	app.Router.Handle("/create", m.Authenticate(m.RequireAuthentication(http.HandlerFunc(emp.CreateVacancy))))
	app.Router.Handle("/delete", m.Authenticate(m.RequireAuthentication(http.HandlerFunc(emp.DeleteVacancy))))
	app.Router.Handle("/update", m.Authenticate(m.RequireAuthentication(http.HandlerFunc(emp.UpdateVacancy))))

}
