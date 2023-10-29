package service

import (
	"Basic-Enrollment-System/model"
	"Basic-Enrollment-System/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseService interface {
	GetAllCourseData() ([]model.Course, error)
	GetOneCourseData(string) (*model.Course, error)
	CreateCourseData(model.Course) (*model.Course, error)
	EditCourseData(string, *model.Course) (*model.Course, error)
	DeleteCourseData(string) (*model.Course, error)
}

type courseService struct {
	courseRepository repository.CourseRepository
}

// NewCourseService -> returns new Course service
func NewCourseService(r repository.CourseRepository) courseService {
	return courseService{
		courseRepository: r,
	}
}

func (r courseService) GetAllCourseData() ([]model.Course, error) {
	return r.courseRepository.GetAllCourseData()
}

func (r courseService) GetOneCourseData(id string) (*model.Course, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.courseRepository.GetOneCourseData(objectId)
}

func (r courseService) CreateCourseData(patient model.Course) (*model.Course, error) {
	result, err := r.courseRepository.CreateCourseData(patient)
	return result, err
}

func (r courseService) EditCourseData(id string, patient *model.Course) (*model.Course, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.courseRepository.EditCourseData(objectId, patient)
}

func (r courseService) DeleteCourseData(id string) (*model.Course, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.courseRepository.DeleteCourseData(objectId)
}
