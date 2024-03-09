package main

import (
	"awesomeProject/internal/db"
	"awesomeProject/pkg/controllers"
	"awesomeProject/pkg/handlers"
	"awesomeProject/pkg/repositories"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	err := godotenv.Load() // Load .env variables
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	connectDatabase := db.ConnectDatabase()
	mux := http.NewServeMux()

	authorizationHandler := makeAuthorizationHandler(connectDatabase)

	mux.HandleFunc("POST /teacher/login", authorizationHandler.TeacherLogin)
	mux.HandleFunc("POST /student/login", authorizationHandler.StudentLogin)
	mux.HandleFunc("POST /administrator/login", authorizationHandler.AdministratorLogin)

	studentHandler := makeStudentHandler(connectDatabase)

	mux.HandleFunc("GET /students", authorizationHandler.AuthMiddleware(studentHandler.GetAllStudents))
	mux.HandleFunc("POST /students", authorizationHandler.AuthMiddleware(studentHandler.CreateStudent))

	teacherHandler := makeTeacherHandler(connectDatabase)

	mux.HandleFunc("GET /teachers", authorizationHandler.AuthMiddleware(teacherHandler.GetAllTeachers))
	mux.HandleFunc("POST /teachers", authorizationHandler.AuthMiddleware(teacherHandler.CreateTeacher))

	courseHandler := makeCourseHandler(connectDatabase)

	mux.HandleFunc("POST /courses", authorizationHandler.AuthMiddleware(courseHandler.CreateCourse))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Failed to start server")
		panic(err)
	}
	fmt.Println("Server is running on port 8080")
}

func makeAuthorizationHandler(db *gorm.DB) *handlers.AuthorizationHandler {
	repository := repositories.NewAuthorizationRepository(db)
	controller := controllers.NewAuthorizationController(repository)
	return handlers.NewAuthorizationHandler(controller)
}

func makeStudentHandler(db *gorm.DB) *handlers.StudentsHandler {
	repository := repositories.NewStudentRepository(db)
	controller := controllers.NewStudentController(repository)
	return handlers.NewStudentsHandler(controller)
}

func makeTeacherHandler(db *gorm.DB) *handlers.TeacherHandler {
	repository := repositories.NewTeacherRepository(db)
	controller := controllers.NewTeacherController(repository)
	return handlers.NewTeacherHandler(controller)
}

func makeCourseHandler(db *gorm.DB) *handlers.CourseHandler {
	repository := repositories.NewCourseRepository(db)
	controller := controllers.NewCourseController(repository)
	return handlers.NewCourseHandler(controller)
}
