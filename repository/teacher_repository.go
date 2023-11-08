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

type TeacherRepository interface {
	GetAllTeacherData() ([]model.Teacher, error)
	GetOneTeacherData(primitive.ObjectID) (*model.Teacher, error)
	CreateTeacherData(model.Teacher) (*model.Teacher, error)
	EditTeacherData(primitive.ObjectID, *model.Teacher) (*model.Teacher, error)
	DeleteTeacherData(primitive.ObjectID) (*model.Teacher, error)
	GetAllTeacherByFacultyName(string) ([]model.Teacher, error)
}

type teacherRepository struct {
	mongoClient       *mongo.Client
	teacherCollection *mongo.Collection
}

func NewTeacherRepository(mongoClient *mongo.Client) teacherRepository {
	return teacherRepository{
		mongoClient:       mongoClient,
		teacherCollection: config.GetCollection(mongoClient, "teacher"),
	}
}

func (r teacherRepository) GetAllTeacherData() ([]model.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := r.teacherCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)

	var dataList []model.Teacher
	for results.Next(ctx) {
		var data model.Teacher
		if err = results.Decode(&data); err != nil {
			return nil, err
		}

		dataList = append(dataList, data)
	}
	return dataList, nil
}

func (r teacherRepository) GetOneTeacherData(id primitive.ObjectID) (*model.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data *model.Teacher

	filter := bson.D{{Key: "_id", Value: id}}

	err := r.teacherCollection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r teacherRepository) getTeacherIdNextValue() int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "seq", Value: 1}}}}
	idQuery := bson.D{{Key: "_id", Value: "couter"}}
	after := options.After
	option := options.FindOneAndUpdateOptions{ReturnDocument: &after}

	var result model.Couter
	err := r.teacherCollection.FindOneAndUpdate(ctx, idQuery, update, &option).Decode(&result)

	if err != nil {
		r.teacherCollection.InsertOne(ctx, bson.M{"_id": "couter", "seq": 1})
		fmt.Println("yung error na kub")
		return 1
	}

	return result.Seq
}

func (r teacherRepository) CreateTeacherData(data model.Teacher) (*model.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	couterIdInt := r.getTeacherIdNextValue()
	dataId := strconv.Itoa(couterIdInt)
	data.TeacherId = dataId

	result, err := r.teacherCollection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	var newData *model.Teacher
	err = r.teacherCollection.FindOne(ctx, bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&newData)
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r teacherRepository) EditTeacherData(id primitive.ObjectID, data *model.Teacher) (*model.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: data}}

	var result model.Teacher
	err := r.teacherCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r teacherRepository) DeleteTeacherData(id primitive.ObjectID) (*model.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}

	var result model.Teacher
	err := r.teacherCollection.FindOneAndDelete(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r teacherRepository) GetAllTeacherByFacultyName(facultyName string) ([]model.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "faculty", Value: facultyName}}

	cursor, err := r.teacherCollection.Find(ctx, filter)
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
