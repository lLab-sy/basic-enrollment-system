package repository

import (
	"Basic-Enrollment-System/config"
	"Basic-Enrollment-System/model"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentRepository interface {
	GetAllStudentData() ([]model.Student, error)
	GetOneStudentData(primitive.ObjectID) (*model.Student, error)
	CreateStudentData(model.Student) (*model.Student, error)
	EditStudentData(primitive.ObjectID, *model.Student) (*model.Student, error)
	DeleteStudentData(primitive.ObjectID) (*model.Student, error)
	GetAllByFacultyName(string) ([]model.Teacher, error)
}

type studentRepository struct {
	mongoClient       *mongo.Client
	studentCollection *mongo.Collection
}

func NewStudentRepository(mongoClient *mongo.Client) studentRepository {
	return studentRepository{
		mongoClient:       mongoClient,
		studentCollection: config.GetCollection(mongoClient, "student"),
	}
}

func (r studentRepository) GetAllStudentData() ([]model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := r.studentCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)

	var dataList []model.Student
	for results.Next(ctx) {
		var data model.Student
		if err = results.Decode(&data); err != nil {
			return nil, err
		}

		dataList = append(dataList, data)
	}
	return dataList, nil
}

func (r studentRepository) GetOneStudentData(id primitive.ObjectID) (*model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data *model.Student

	filter := bson.D{{Key: "_id", Value: id}}

	err := r.studentCollection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r studentRepository) getStudentIdNextValue() int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "seq", Value: 1}}}}
	idQuery := bson.D{{Key: "_id", Value: "couter"}}
	after := options.After
	option := options.FindOneAndUpdateOptions{ReturnDocument: &after}

	var result model.Couter
	err := r.studentCollection.FindOneAndUpdate(ctx, idQuery, update, &option).Decode(&result)

	if err != nil {
		r.studentCollection.InsertOne(ctx, bson.M{"_id": "couter", "seq": 1})
		fmt.Println("yung error na kub")
		return 1
	}

	return result.Seq
}

func (r studentRepository) CreateStudentData(data model.Student) (*model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	couterIdInt := r.getStudentIdNextValue()
	dataId := strconv.Itoa(couterIdInt)
	data.StudentId = dataId

	result, err := r.studentCollection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	var newData *model.Student
	err = r.studentCollection.FindOne(ctx, bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&newData)
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r studentRepository) EditStudentData(id primitive.ObjectID, data *model.Student) (*model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: data}}

	var result model.Student
	err := r.studentCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r studentRepository) DeleteStudentData(id primitive.ObjectID) (*model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}

	var result model.Student
	err := r.studentCollection.FindOneAndDelete(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r studentRepository) GetAllByFacultyName(facultyName string) ([]model.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "faculty", Value: facultyName}}

	cursor, err := r.studentCollection.Find(ctx, filter)
	if err != nil {
		return []model.Teacher{}, err
	}
	defer cursor.Close(ctx)

	// Iterate over the results
	var results []model.Teacher
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	return results, nil
}
