package handlers

import (
	"awesomeProject/pkg/controllers"
	"awesomeProject/pkg/models"
	"encoding/json"
	"net/http"
)

type TeacherHandler struct {
	teacherController *controllers.TeacherController
}

func NewTeacherHandler(teacherController *controllers.TeacherController) *TeacherHandler {
	return &TeacherHandler{teacherController}
}

func (h *TeacherHandler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	if requestContext.Role != models.AdministratorRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var teacher models.Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		http.Error(w, "Error decoding teacher creation request", http.StatusBadRequest)
		return
	}
	err = h.teacherController.CreateTeacher(&teacher)
	if err != nil {
		http.Error(w, "Error creating teacher", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TeacherHandler) GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	if requestContext.Role != models.AdministratorRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	teachers, err := h.teacherController.GetAllTeachers()
	if err != nil {
		http.Error(w, "Error getting teachers", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(teachers)
	if err != nil {
		http.Error(w, "Error encoding teachers", http.StatusInternalServerError)
		return
	}
}
