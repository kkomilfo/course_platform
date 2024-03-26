package controllers

import (
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/repositories"
)

type StudentController struct {
	studentRepository *repositories.StudentRepository
}

type StudentWorkRequest struct {
	SubjectID uint `json:"subject_id"`
	Files     []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"files"`
}

type CommentRequest struct {
	Content string `json:"content"`
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

func (c *StudentController) CreateStudentWork(studentID uint, request StudentWorkRequest) error {
	studentWorkFiles := make([]models.StudentWorkFile, len(request.Files))
	for i, file := range request.Files {
		studentWorkFiles[i] = models.StudentWorkFile{
			Name: file.Name,
			URL:  file.URL,
		}

	}
	work := models.StudentWork{
		SubjectID: request.SubjectID,
		StudentID: studentID,
		Files:     studentWorkFiles,
	}
	return c.studentRepository.CreateStudentWork(&work)
}

func (c *StudentController) Comment(workID uint, userID uint, role models.Role, comment CommentRequest) error {
	var stringRole string
	if role == models.StudentRole {
		stringRole = "student"
	} else {
		stringRole = "teacher"
	}
	var comment1 = models.Comment{
		UserID:        userID,
		UserType:      stringRole,
		StudentWorkID: workID,
		Content:       comment.Content,
	}
	return c.studentRepository.Comment(&comment1)
}
