package pkg

type EmployerView struct {
	ID          string
	Username    string
	Password    string
	CompanyName string
	Vacancies   []VacancyView
}

type VacancyView struct {
	ID          string
	Name        string
	Description string
}
