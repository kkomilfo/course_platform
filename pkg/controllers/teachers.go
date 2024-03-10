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

type GradeRequest struct {
	StudentWorkID uint `json:"student_work_id"`
	Grade         int  `json:"grade"`
}

func (c *TeacherController) GradeStudentWork(request *GradeRequest) error {
	return c.teacherRepository.GradeStudentWork(request.StudentWorkID, request.Grade)
}
