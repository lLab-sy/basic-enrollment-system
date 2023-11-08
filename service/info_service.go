package service

import (
	"Basic-Enrollment-System/model"
	"Basic-Enrollment-System/repository"
)

type InfoService interface {
	GetTeacherAndStudentByFaculty(string) (model.FacultyMember, error)
}

type infoService struct {
	studentRepository repository.StudentRepository
	teacherRepository repository.TeacherRepository
}

// NewInfoService -> returns new Info service
func NewInfoService(rs repository.StudentRepository, rt repository.TeacherRepository) infoService {
	return infoService{
		studentRepository: rs,
		teacherRepository: rt,
	}
}

func (r infoService) GetTeacherAndStudentByFaculty(facultyName string) (model.FacultyMember, error) {
	var facultyMember model.FacultyMember

	facultyMember.FacultyName = facultyName

	students, err := r.studentRepository.GetAllStudentByFacultyName(facultyName)
	if err != nil {
		return model.FacultyMember{}, nil
	}
	facultyMember.Students = students

	teachers, err := r.teacherRepository.GetAllTeacherByFacultyName(facultyName)
	if err != nil {
		return model.FacultyMember{}, nil
	}
	facultyMember.Teachers = teachers

	return facultyMember, nil
}
