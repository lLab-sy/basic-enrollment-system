package service

import (
	"Basic-Enrollment-System/model"
	"Basic-Enrollment-System/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeacherService interface {
	GetAllTeacherData() ([]model.Teacher, error)
	GetOneTeacherData(string) (*model.Teacher, error)
	CreateTeacherData(model.Teacher) (*model.Teacher, error)
	EditTeacherData(string, *model.Teacher) (*model.Teacher, error)
	DeleteTeacherData(string) (*model.Teacher, error)
}

type teacherService struct {
	teacherRepository repository.TeacherRepository
}

// NewTeacherService -> returns new Teacher service
func NewTeacherService(r repository.TeacherRepository) teacherService {
	return teacherService{
		teacherRepository: r,
	}
}

func (r teacherService) GetAllTeacherData() ([]model.Teacher, error) {
	return r.teacherRepository.GetAllTeacherData()
}

func (r teacherService) GetOneTeacherData(id string) (*model.Teacher, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.teacherRepository.GetOneTeacherData(objectId)
}

func (r teacherService) CreateTeacherData(patient model.Teacher) (*model.Teacher, error) {
	result, err := r.teacherRepository.CreateTeacherData(patient)
	return result, err
}

func (r teacherService) EditTeacherData(id string, patient *model.Teacher) (*model.Teacher, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.teacherRepository.EditTeacherData(objectId, patient)
}

func (r teacherService) DeleteTeacherData(id string) (*model.Teacher, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.teacherRepository.DeleteTeacherData(objectId)
}
