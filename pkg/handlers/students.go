package handlers

import (
	"awesomeProject/pkg/controllers"
	"awesomeProject/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type StudentsHandler struct {
	studentController *controllers.StudentController
}

func NewStudentsHandler(studentController *controllers.StudentController) *StudentsHandler {
	return &StudentsHandler{studentController}
}

func (h *StudentsHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	if requestContext.Role != models.AdministratorRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Error decoding student creation request", http.StatusBadRequest)
		return
	}
	err = h.studentController.CreateStudent(&student)
	if err != nil {
		http.Error(w, "Error creating student", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *StudentsHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	if requestContext.Role == models.StudentRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	students, err := h.studentController.GetAllStudents()
	if err != nil {
		http.Error(w, "Error getting students", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		http.Error(w, "Error encoding students", http.StatusInternalServerError)
		return
	}
}

func (h *StudentsHandler) CreateStudentWork(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	if requestContext.Role != models.StudentRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var studentWorkRequest controllers.StudentWorkRequest
	err := json.NewDecoder(r.Body).Decode(&studentWorkRequest)
	if err != nil {
		http.Error(w, "Error decoding student work creation request", http.StatusBadRequest)
		return

	}

	fmt.Println(studentWorkRequest)
	err = h.studentController.CreateStudentWork(requestContext.UserID, studentWorkRequest)
	if err != nil {
		http.Error(w, "Error creating student work", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *StudentsHandler) Comment(writer http.ResponseWriter, request *http.Request) {
	requestContext := request.Context().Value(RequestContextKey).(RequestContext)
	var comment controllers.CommentRequest
	err := json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		http.Error(writer, "Error decoding comment request", http.StatusBadRequest)
		return
	}
	courseIDString := request.PathValue("id")
	workID, _ := strconv.ParseUint(courseIDString, 10, 64)
	err = h.studentController.Comment(uint(workID), requestContext.UserID, requestContext.Role, comment)
	if err != nil {
		http.Error(writer, "Error commenting", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
