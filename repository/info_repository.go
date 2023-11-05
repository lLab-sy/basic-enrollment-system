package repository

import (
	"Basic-Enrollment-System/config"
	"Basic-Enrollment-System/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InfoRepository interface {
	GetAllMemberInFaculty(string) (model.FacultyMember, error)
}

type infoRepository struct {
	mongoClient       *mongo.Client
	studentCollection *mongo.Collection
	teacherCollection *mongo.Collection
}

func NewInfoRepository(mongoClient *mongo.Client) infoRepository {
	return infoRepository{
		mongoClient:       mongoClient,
		studentCollection: config.GetCollection(mongoClient, "student"),
		teacherCollection: config.GetCollection(mongoClient, "teacher"),
	}
}

func (r infoRepository) GetAllMemberInFaculty(facultyName string) (model.FacultyMember, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := r.infoCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)

	var dataList []model.Info
	for results.Next(ctx) {
		var data model.Info
		if err = results.Decode(&data); err != nil {
			return nil, err
		}

		dataList = append(dataList, data)
	}
	return dataList, nil
}
