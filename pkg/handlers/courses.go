package handlers

import (
	"awesomeProject/pkg/controllers"
	"awesomeProject/pkg/models"
	"encoding/json"
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
