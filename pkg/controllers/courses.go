package controllers

import (
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/repositories"
	"time"
)

type EnrollStudentRequest struct {
	StudentID uint `json:"student_id"`
	CourseID  uint `json:"course_id"`
}

type ModuleRequest struct {
	Title string `json:"title"`
}

type SubjectRequest struct {
	Title       string               `json:"title"`
	Description string               `json:"description"`
	DueDate     time.Time            `json:"due_date"`
	Type        string               `json:"type"`
	Files       []SubjectFileRequest `json:"files"`
}

type SubjectFileRequest struct {
	URL  string `json:"url"`
	Name string `json:"name"`
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

func (c *CourseController) AddModuleToCourse(courseID uint, module *ModuleRequest) error {
	model := models.Module{Title: module.Title}
	return c.repository.AddModuleToCourse(courseID, &model)
}

func (c *CourseController) AddSubjectToModule(moduleID uint, subject *SubjectRequest) error {
	files := make([]models.File, 0)
	for _, file := range subject.Files {
		files = append(files, models.File{URL: file.URL, Name: file.Name})
	}
	model := models.Subject{
		Title:       subject.Title,
		Description: subject.Description,
		DueDate:     subject.DueDate,
		Type:        models.SubjectType(subject.Type),
		Files:       files,
	}
	return c.repository.AddSubjectToModule(moduleID, &model)
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
