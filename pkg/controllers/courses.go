package controllers

import (
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/repositories"
	"fmt"
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

func (c *CourseController) GetCourseByID(courseID uint) (CourseDetailsResponse, error) {
	course, err := c.repository.FindCourseByID(courseID)
	if err != nil {
		return CourseDetailsResponse{}, err
	}
	return CourseDetailsResponseFromModel(course), nil
}

func (c *CourseController) GetAllCoursesByStudentID(studentID uint) ([]StudentCourseResponse, error) {
	courses, err := c.repository.FindAllByStudentID(studentID)

	if err != nil {
		return nil, err
	}
	var courseResponses []StudentCourseResponse
	for _, course := range courses {
		courseResponses = append(courseResponses, StudentCourseResponseFromModel(course))
	}
	return courseResponses, nil
}

type StudentCourseResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func StudentCourseResponseFromModel(course models.Course) StudentCourseResponse {
	return StudentCourseResponse{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		ImageURL:    course.ImageURL,
	}
}

type CourseDetailsResponse struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	ImageURL    string           `json:"image_url"`
	Modules     []ModuleResponse `json:"modules"`
}

type ModuleResponse struct {
	ID       uint              `json:"id"`
	Title    string            `json:"title"`
	Subjects []SubjectResponse `json:"subjects"`
}

type SubjectResponse struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	DueDate     string                `json:"due_date"`
	Type        string                `json:"type"`
	Files       []SubjectFileResponse `json:"files"`
}

type SubjectFileResponse struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func CourseDetailsResponseFromModel(course models.Course) CourseDetailsResponse {
	var modules []ModuleResponse
	for _, module := range course.Modules {
		var subjects []SubjectResponse
		for _, subject := range module.Subjects {
			var files []SubjectFileResponse
			for _, file := range subject.Files {
				files = append(files, SubjectFileResponse{
					URL:  file.URL,
					Name: file.Name,
				})
			}
			subjects = append(subjects, SubjectResponse{
				ID:          subject.ID,
				Title:       subject.Title,
				Description: subject.Description,
				DueDate:     subject.DueDate.Format(time.RFC3339),
				Type:        string(subject.Type),
				Files:       files,
			})
		}
		modules = append(modules, ModuleResponse{
			ID:       module.ID,
			Title:    module.Title,
			Subjects: subjects,
		})
	}
	return CourseDetailsResponse{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		ImageURL:    course.ImageURL,
		Modules:     modules,
	}
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

func (c *CourseController) GetSubjectTaskForStudent(subjectID uint, studentID uint) (SubjectTaskResponse, error) {
	subject, err := c.repository.FindSubject(subjectID)
	if err != nil {
		return SubjectTaskResponse{}, err
	}
	work, err := c.repository.FindStudentWork(subjectID, studentID)
	if err != nil {
		return SubjectTaskResponse{}, err
	}
	return SubjectTaskResponseFromModel(subject, work), nil
}

type EntrolledStudentResponse struct {
	ID        uint   `json:"id"`
	FullName  string `json:"full_name"`
	AvatarURL string `json:"avatar_url"`
	Works     []Work `json:"student_works"`
}

type Work struct {
	SubjectID uint  `json:"subject_id"`
	Grade     *uint `json:"grade"`
}

func (c *CourseController) GetAllEntrolledStudentsByCourseID(id uint) ([]EntrolledStudentResponse, error) {
	course, err := c.repository.FindCourseByID(id)

	var tasks []uint
	for _, module := range course.Modules {
		if module.Subjects == nil {
			continue
		}
		for _, subject := range module.Subjects {
			if subject.Type == "Task" {
				tasks = append(tasks, subject.ID)
			}
		}
	}

	students, err := c.repository.FindAllEntrolledStudentsByCourseID(id)
	if err != nil {
		return nil, err
	}
	var studentResponses []EntrolledStudentResponse
	for _, student := range students {
		works, err := c.repository.GetStudentsWorksWithSubject(student.ID, tasks)

		if err != nil {
			return nil, err
		}

		// Create a map to track Works based on SubjectID
		workMap := make(map[uint]Work)
		for _, work := range works {
			fmt.Println(work.SubjectID)
			workMap[work.SubjectID] = Work{
				SubjectID: work.SubjectID,
				Grade:     work.Grade,
			}
		}

		// Initialize empty Works array
		result := make([]Work, 0, len(tasks)) // Pre-allocate space based on the number of tasks

		// Create Work objects for all tasks, filling grades from workMap
		for _, task := range tasks {
			work, exists := workMap[task] // Adjust task.SubjectID based on your tasks data
			work.SubjectID = task
			if !exists {
				work.Grade = nil // Set grade to nil if work doesn't exist
			}
			result = append(result, work)
		}

		studentResponses = append(studentResponses, EntrolledStudentResponse{
			ID:        student.ID,
			FullName:  student.FullName,
			AvatarURL: student.AvatarURL,
			Works:     result,
		})
	}
	return studentResponses, nil
}

type SubjectTaskResponse struct {
	ID            uint                         `json:"id"`
	StudentWorkID uint                         `json:"student_work_id"`
	Title         string                       `json:"title"`
	Description   string                       `json:"description"`
	DueDate       string                       `json:"due_date"`
	Comments      []SubjectTaskCommentResponse `json:"comments"`
	TaskFiles     []SubjectFileResponse        `json:"task_files"`
	StudentFiles  []SubjectFileResponse        `json:"student_files"`
	Grade         *uint                        `json:"grade"`
}

type SubjectTaskCommentResponse struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	UserType string `json:"user_type"`
}

func SubjectTaskResponseFromModel(subject models.Subject, work models.StudentWork) SubjectTaskResponse {
	var files []SubjectFileResponse
	for _, file := range subject.Files {
		files = append(files, SubjectFileResponse{
			URL:  file.URL,
			Name: file.Name,
		})
	}
	var comments []SubjectTaskCommentResponse
	for _, comment := range work.Comments {
		comments = append(comments, SubjectTaskCommentResponse{
			ID:       comment.ID,
			Content:  comment.Content,
			UserType: comment.UserType,
		})
	}
	var studentFiles []SubjectFileResponse
	for _, file := range work.Files {
		studentFiles = append(studentFiles, SubjectFileResponse{
			URL:  file.URL,
			Name: file.Name,
		})
	}
	return SubjectTaskResponse{
		ID:            subject.ID,
		StudentWorkID: work.ID,
		Title:         subject.Title,
		Description:   subject.Description,
		DueDate:       subject.DueDate.Format(time.RFC3339),
		Comments:      comments,
		TaskFiles:     files,
		StudentFiles:  studentFiles,
		Grade:         work.Grade,
	}
}
