package controllers

import (
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/repositories"
)

type EnrollStudentRequest struct {
	StudentID uint `json:"student_id"`
	CourseID  uint `json:"course_id"`
}

type CourseController struct {
	repository *repositories.CourseRepository
}

func NewCourseController(repository *repositories.CourseRepository) *CourseController {
	return &CourseController{repository}
}

func (c *CourseController) CreateCourse(course *models.Course) error {
	return c.repository.Create(course)
}

func (c *CourseController) EnrollStudent(request *EnrollStudentRequest) error {
	return c.repository.EnrollStudent(request.StudentID, request.CourseID)
}

func (c *CourseController) GetAllCoursesByTeacherID(teacherID uint) ([]CourseResponse, error) {
	courses, err := c.repository.FindAllByTeacherID(teacherID)
	if err != nil {
		return nil, err
	}
	var courseResponses []CourseResponse
	for _, course := range courses {
		courseResponses = append(courseResponses, CourseResponseFromModel(course))
	}
	return courseResponses, nil
}

type CourseResponse struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	ImageURL    string            `json:"image_url"`
	Students    []StudentResponse `json:"students"`
}

type StudentResponse struct {
	ID        uint   `json:"id"`
	FullName  string `json:"full_name"`
	AvatarURL string `json:"avatar_url"`
}

func CourseResponseFromModel(course models.Course) CourseResponse {
	var students []StudentResponse
	for _, student := range course.Students {
		students = append(students, StudentResponse{
			ID:        student.ID,
			FullName:  student.FullName,
			AvatarURL: student.AvatarURL,
		})
	}
	return CourseResponse{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		ImageURL:    course.ImageURL,
		Students:    students,
	}
}
