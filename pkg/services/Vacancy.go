package services

import (
	"context"
	"log"
	"main/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VacancyService struct {
	client *mongo.Database
}

func NewVacancyService(client *mongo.Database) *VacancyService {
	return &VacancyService{client: client}
}

func (s *VacancyService) Create(obj models.Vacancy) (models.Vacancy, error) {
	collection := s.client.Collection("vacancies")
	obj.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(context.TODO(), obj)
	if err != nil {
		return models.Vacancy{}, err
	}
	id := res.InsertedID.(primitive.ObjectID)

	filter := bson.M{"_id": id}
	var createdVacancy models.Vacancy
	err = collection.FindOne(context.TODO(), filter).Decode(&createdVacancy)
	if err != nil {
		log.Printf("Ошибка при получении созданного документа: %v", err)
		return models.Vacancy{}, err
	}

	return createdVacancy, nil
}

func (s *VacancyService) Update(id primitive.ObjectID, updatedVacancy models.Vacancy) error {
	collection := s.client.Collection("vacancies")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedVacancy}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("Ошибка при обновлении документа: %v", err)
		return err
	}

	return nil
}

func (s *VacancyService) Delete(id primitive.ObjectID) error {
	collection := s.client.Collection("vacancies")

	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("Ошибка при удалении документа: %v", err)
		return err
	}

	return nil
}

func (s *VacancyService) Get(id primitive.ObjectID) (models.Vacancy, error) {
	var vacancy models.Vacancy

	collection := s.client.Collection("vacancies")

	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&vacancy)
	if err != nil {
		log.Printf("Ошибка при получении документа: %v", err)
		return models.Vacancy{}, err
	}

	return vacancy, nil
}

func (s *VacancyService) GetAll() ([]models.Vacancy, error) {
	var vacancies []models.Vacancy

	collection := s.client.Collection("vacancies")

	// Выполняем запрос ко всем документам в коллекции
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Итерируемся по результатам запроса
	for cursor.Next(context.TODO()) {
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
