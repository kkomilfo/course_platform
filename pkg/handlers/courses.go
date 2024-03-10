package handlers

import (
	"awesomeProject/pkg/controllers"
	"awesomeProject/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type CourseHandler struct {
	controller *controllers.CourseController
}

func NewCourseHandler(courseController *controllers.CourseController) *CourseHandler {
	return &CourseHandler{courseController}
}

func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	if requestContext.Role != models.AdministratorRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.controller.CreateCourse(&course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CourseHandler) EnrollStudent(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	fmt.Println(requestContext.Role)
	if requestContext.Role != models.TeacherRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var enrollment controllers.EnrollStudentRequest
	err := json.NewDecoder(r.Body).Decode(&enrollment)
	if err != nil {
		http.Error(w, "Error decoding enrollment request", http.StatusBadRequest)
		return
	}
	err = h.controller.EnrollStudent(&enrollment)
	if err != nil {
		http.Error(w, "Error enrolling student", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CourseHandler) GetAllCoursesByTeacherID(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	if requestContext.Role != models.TeacherRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	teacherID := requestContext.UserID
	courses, err := h.controller.GetAllCoursesByTeacherID(teacherID)
	if err != nil {
		http.Error(w, "Error getting courses", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(courses)
	if err != nil {
		http.Error(w, "Error encoding courses", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
