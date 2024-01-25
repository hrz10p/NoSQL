package services

import (
	"context"
	"log"
	"main/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VacancyService struct {
	client *mongo.Database
}

func NewVacancyService(client *mongo.Database) *VacancyService {
	return &VacancyService{client: client}
}

func (s *VacancyService) create(obj models.Vacancy) {
	collection := s.client.Collection("vacancies")

	_, err := collection.InsertOne(context.TODO(), obj)
	if err != nil {
		log.Fatal(err)
	}

}

func (s *VacancyService) update() {

}

func (s *VacancyService) delete() {

}

func (s *VacancyService) get() {

}

func (s *VacancyService) GetAll() ([]models.Vacancy, error) {
	var vacancies []models.Vacancy

	// Создаем контекст для выполнения операции ввода/вывода
	ctx := context.TODO()

	// Выполняем запрос ко всем документам в коллекции
	cursor, err := s.client.Collection("vacancies").Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Итерируемся по результатам запроса
	for cursor.Next(ctx) {
		var vacancy models.Vacancy
		if err := cursor.Decode(&vacancy); err != nil {
			log.Printf("Ошибка при декодировании документа: %v", err)
			return nil, err
		}

		// Добавляем декодированный документ в массив
		vacancies = append(vacancies, vacancy)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Ошибка при обработке результатов запроса: %v", err)
		return nil, err
	}

	return vacancies, nil
}
