package controllers

import (
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/repositories"
)

type CourseController struct {
	repository *repositories.CourseRepository
}

func NewCourseController(repository *repositories.CourseRepository) *CourseController {
	return &CourseController{repository}
}

func (c *CourseController) CreateCourse(course *models.Course) error {
	return c.repository.Create(course)
}
