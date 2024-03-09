package controllers

import (
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/repositories"
)

type TeacherController struct {
	teacherRepository *repositories.TeacherRepository
}

func NewTeacherController(teacherRepository *repositories.TeacherRepository) *TeacherController {
	return &TeacherController{teacherRepository}
}

func (c *TeacherController) CreateTeacher(teacher *models.Teacher) error {
	return c.teacherRepository.Create(teacher)
}

func (c *TeacherController) GetAllTeachers() ([]models.Teacher, error) {
	return c.teacherRepository.FindAll()
}
