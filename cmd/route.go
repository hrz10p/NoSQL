package main

import (
	"fmt"
	"main/pkg/logger"
)

// InitializeRoutes sets up the application routes
func (app *Application) InitializeRoutes() {
	obj, err := app.Service.VacancyService.GetAll()
	if err != nil {
		logger.GetLogger().Error(err.Error())
	}

	fmt.Println(obj)

}
