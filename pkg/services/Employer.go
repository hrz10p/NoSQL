package services

import (
	"context"
	"log"
	"main/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployerService struct {
	client *mongo.Database
}

func NewEmployerService(client *mongo.Database) *EmployerService {
	return &EmployerService{client: client}
}

func (s *EmployerService) Create(employer models.Employer) error {
	collection := s.client.Collection("employers")

	_, err := collection.InsertOne(context.TODO(), employer)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployerService) Update(id primitive.ObjectID, updatedEmployer models.Employer) error {
	collection := s.client.Collection("employers")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedEmployer}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("Ошибка при обновлении документа: %v", err)
		return err
	}

	return nil
}

func (s *EmployerService) Delete(id primitive.ObjectID) error {
	collection := s.client.Collection("employers")

	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("Ошибка при удалении документа: %v", err)
		return err
	}

	return nil
}

func (s *EmployerService) Get(id primitive.ObjectID) (models.Employer, error) {
	var employer models.Employer

	collection := s.client.Collection("employers")

	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&employer)
	if err != nil {
		log.Printf("Ошибка при получении документа: %v", err)
		return models.Employer{}, err
	}

	return employer, nil
}

func (s *EmployerService) GetByUsername(username string) (models.Employer, error) {
	var employer models.Employer

	collection := s.client.Collection("employers")

	filter := bson.M{"username": username}
	err := collection.FindOne(context.TODO(), filter).Decode(&employer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Employer{}, models.ErrNoRecord
		}
		log.Printf("Ошибка при получении документа: %v", err)
		return models.Employer{}, err
	}

	return employer, nil
}

func (s *EmployerService) GetAll() ([]models.Employer, error) {
	var employers []models.Employer

	collection := s.client.Collection("employers")

	// Выполняем запрос ко всем документам в коллекции
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Итерируемся по результатам запроса
	for cursor.Next(context.TODO()) {
		var employer models.Employer
		if err := cursor.Decode(&employer); err != nil {
			log.Printf("Ошибка при декодировании документа: %v", err)
			return nil, err
		}

		// Добавляем декодированный документ в массив
		employers = append(employers, employer)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Ошибка при обработке результатов запроса: %v", err)
		return nil, err
	}

	return employers, nil
}
