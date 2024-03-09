package handlers

import (
	"awesomeProject/pkg/controllers"
	"awesomeProject/pkg/models"
	"encoding/json"
	"net/http"
)

type StudentsHandler struct {
	studentController *controllers.StudentController
}

func NewStudentsHandler(studentController *controllers.StudentController) *StudentsHandler {
	return &StudentsHandler{studentController}
}

func (h *StudentsHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
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
	if requestContext.Role != models.AdministratorRole {
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
