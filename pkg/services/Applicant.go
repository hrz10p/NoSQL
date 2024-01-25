package services

import (
	"context"
	"log"
	"main/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplicantService struct {
	client *mongo.Database
}

func NewApplicantService(client *mongo.Database) *ApplicantService {
	return &ApplicantService{client: client}
}

func (s *ApplicantService) Create(applicant models.Applicant) error {
	collection := s.client.Collection("applicants")

	_, err := collection.InsertOne(context.TODO(), applicant)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (s *ApplicantService) Update(id primitive.ObjectID, updatedApplicant models.Applicant) error {
	collection := s.client.Collection("applicants")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedApplicant}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("Ошибка при обновлении документа: %v", err)
		return err
	}

	return nil
}

func (s *ApplicantService) Delete(id primitive.ObjectID) error {
	collection := s.client.Collection("applicants")

	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("Ошибка при удалении документа: %v", err)
		return err
	}

	return nil
}

func (s *ApplicantService) Get(id primitive.ObjectID) (models.Applicant, error) {
	var applicant models.Applicant

	collection := s.client.Collection("applicants")

	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&applicant)
	if err != nil {
		log.Printf("Ошибка при получении документа: %v", err)
		return models.Applicant{}, err
	}

	return applicant, nil
}

func (s *ApplicantService) GetByUsername(username string) (models.Applicant, error) {
	var applicant models.Applicant

	collection := s.client.Collection("applicants")

	filter := bson.M{"username": username}

	err := collection.FindOne(context.TODO(), filter).Decode(&applicant)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Applicant{}, models.ErrNoRecord
		}
		log.Printf("Ошибка при получении документа: %v", err)
		return models.Applicant{}, err
	}

	return applicant, nil
}

func (s *ApplicantService) GetAll() ([]models.Applicant, error) {
	var applicants []models.Applicant

	collection := s.client.Collection("applicants")

	// Выполняем запрос ко всем документам в коллекции
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Итерируемся по результатам запроса
	for cursor.Next(context.TODO()) {
		var applicant models.Applicant
		if err := cursor.Decode(&applicant); err != nil {
			log.Printf("Ошибка при декодировании документа: %v", err)
			return nil, err
		}

		// Добавляем декодированный документ в массив
		applicants = append(applicants, applicant)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Ошибка при обработке результатов запроса: %v", err)
		return nil, err
	}

	return applicants, nil
}
