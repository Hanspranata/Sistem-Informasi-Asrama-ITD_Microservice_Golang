package repository

import (
	"database/sql"
	"fmt"

	"github.com/username/student-service/internal/student/models"
)

type StudentRepository interface {
	InsertStudent(student *models.Student) error
	UpdateStudent(student *models.Student) error
	DeleteStudent(studentId string) error
	GetStudentById(studentId string) (*models.Student, error)
	GetAllStudents() ([]*models.Student, error)
}

type studentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (s *studentRepository) InsertStudent(student *models.Student) error {
	sql := fmt.Sprintf("INSERT INTO students(id, name, email, room_number) VALUES ('%s', '%s', '%s', '%s')", student.Id, student.Name, student.Email, student.RoomNumber)
	_, err := s.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepository) UpdateStudent(student *models.Student) error {
	sql := fmt.Sprintf("UPDATE students SET name='%s', email='%s', room_number='%s' WHERE id='%s'", student.Name, student.Email, student.RoomNumber, student.Id)
	_, err := s.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepository) DeleteStudent(studentId string) error {
	sql := fmt.Sprintf("DELETE FROM students WHERE id='%s'", studentId)
	_, err := s.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepository) GetStudentById(studentId string) (*models.Student, error) {
	sql := fmt.Sprintf("SELECT id, name, email, room_number FROM students WHERE id='%s'", studentId)
	row := s.db.QueryRow(sql)
	student := &models.Student{}
	err := row.Scan(&student.Id, &student.Name, &student.Email, &student.RoomNumber)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *studentRepository) GetAllStudents() ([]*models.Student, error) {
	sql := "SELECT id, name, email, room_number FROM students"
	rows, err := s.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []*models.Student{}
	for rows.Next() {
		student := &models.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.RoomNumber)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}
