package usecase

import (
	"github.com/username/student-service/internal/student/models"
	"github.com/username/student-service/internal/student/repository"
)

type StudentUsecase interface {
	InsertStudent(student *models.Student) error
	UpdateStudent(student *models.Student) error
	DeleteStudent(studentId string) error
	GetStudentById(studentId string) (*models.Student, error)
	GetAllStudents() ([]*models.Student, error)
}

type studentUsecase struct {
	studentRepo repository.StudentRepository
}

func NewStudentUsecase(studentRepo repository.StudentRepository) StudentUsecase {
	return &studentUsecase{studentRepo}
}

func (s *studentUsecase) InsertStudent(student *models.Student) error {
	err := s.studentRepo.InsertStudent(student)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentUsecase) UpdateStudent(student *models.Student) error {
	err := s.studentRepo.UpdateStudent(student)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentUsecase) DeleteStudent(studentId string) error {
	err := s.studentRepo.DeleteStudent(studentId)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentUsecase) GetStudentById(studentId string) (*models.Student, error) {
	student, err := s.studentRepo.GetStudentById(studentId)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *studentUsecase) GetAllStudents() ([]*models.Student, error) {
	students, err := s.studentRepo.GetAllStudents()
	if err != nil {
		return nil, err
	}
	return students, nil
}
