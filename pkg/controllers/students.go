package controllers

import (
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/repositories"
)

type StudentController struct {
	studentRepository *repositories.StudentRepository
}

func NewStudentController(studentRepository *repositories.StudentRepository) *StudentController {
	return &StudentController{studentRepository}
}

func (c *StudentController) CreateStudent(student *models.Student) error {
	return c.studentRepository.Create(student)
}

func (c *StudentController) GetAllStudents() ([]models.Student, error) {
	return c.studentRepository.FindAll()
}
