package service

import (
	"Basic-Enrollment-System/model"
	"Basic-Enrollment-System/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentService interface {
	GetAllStudentData() ([]model.Student, error)
	GetOneStudentData(string) (*model.Student, error)
	CreateStudentData(model.Student) (*model.Student, error)
	EditStudentData(string, *model.Student) (*model.Student, error)
	DeleteStudentData(string) (*model.Student, error)
}

type studentService struct {
	studentRepository repository.StudentRepository
}

// NewStudentService -> returns new Student service
func NewStudentService(r repository.StudentRepository) studentService {
	return studentService{
		studentRepository: r,
	}
}

func (r studentService) GetAllStudentData() ([]model.Student, error) {
	return r.studentRepository.GetAllStudentData()
}

func (r studentService) GetOneStudentData(id string) (*model.Student, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.studentRepository.GetOneStudentData(objectId)
}

func (r studentService) CreateStudentData(patient model.Student) (*model.Student, error) {
	result, err := r.studentRepository.CreateStudentData(patient)
	return result, err
}

func (r studentService) EditStudentData(id string, patient *model.Student) (*model.Student, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.studentRepository.EditStudentData(objectId, patient)
}

func (r studentService) DeleteStudentData(id string) (*model.Student, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.studentRepository.DeleteStudentData(objectId)
}
