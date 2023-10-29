package repository

import (
	"Basic-Enrollment-System/config"
	"Basic-Enrollment-System/model"
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CourseRepository interface {
	GetAllCourseData() ([]model.Course, error)
	GetOneCourseData(primitive.ObjectID) (*model.Course, error)
	CreateCourseData(model.Course) (*model.Course, error)
	EditCourseData(primitive.ObjectID, *model.Course) (*model.Course, error)
	DeleteCourseData(primitive.ObjectID) (*model.Course, error)
}

type courseRepository struct {
	mongoClient      *mongo.Client
	courseCollection *mongo.Collection
}

func NewCourseRepository(mongoClient *mongo.Client) courseRepository {
	return courseRepository{
		mongoClient:      mongoClient,
		courseCollection: config.GetCollection(mongoClient, "course"),
	}
}

func (r courseRepository) GetAllCourseData() ([]model.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := r.courseCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)

	var dataList []model.Course
	for results.Next(ctx) {
		var data model.Course
		if err = results.Decode(&data); err != nil {
			return nil, err
		}

		dataList = append(dataList, data)
	}
	return dataList, nil
}

func (r courseRepository) GetOneCourseData(id primitive.ObjectID) (*model.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data *model.Course

	filter := bson.D{{Key: "_id", Value: id}}

	err := r.courseCollection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r courseRepository) getCourseIdNextValue() int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "seq", Value: 1}}}}
	idQuery := bson.D{{Key: "_id", Value: "couter"}}
	after := options.After
	option := options.FindOneAndUpdateOptions{ReturnDocument: &after}

	var result model.Couter
	err := r.courseCollection.FindOneAndUpdate(ctx, idQuery, update, &option).Decode(&result)

	if err != nil {
		r.courseCollection.InsertOne(ctx, bson.M{"_id": "couter", "seq": 1})
		fmt.Println("yung error na kub")
		return 1
	}

	return result.Seq
}

func (r courseRepository) CreateCourseData(data model.Course) (*model.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	couterIdInt := r.getCourseIdNextValue()
	dataId := strconv.Itoa(couterIdInt)
	data.CourseId = dataId

	result, err := r.courseCollection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	var newData *model.Course
	err = r.courseCollection.FindOne(ctx, bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&newData)
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r courseRepository) EditCourseData(id primitive.ObjectID, data *model.Course) (*model.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: data}}

	var result model.Course
	err := r.courseCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r courseRepository) DeleteCourseData(id primitive.ObjectID) (*model.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}

	var result model.Course
	err := r.courseCollection.FindOneAndDelete(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
