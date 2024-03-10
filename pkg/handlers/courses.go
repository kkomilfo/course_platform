package handlers

import (
	"awesomeProject/pkg/controllers"
	"awesomeProject/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func (h *CourseHandler) AddModuleToCourse(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	if requestContext.Role != models.TeacherRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	courseIDString := r.PathValue("id")
	courseID, _ := strconv.ParseUint(courseIDString, 10, 64)

	var moduleRequest controllers.ModuleRequest
	err := json.NewDecoder(r.Body).Decode(&moduleRequest)
	if err != nil {
		http.Error(w, "Error decoding module request", http.StatusBadRequest)
		return
	}
	err = h.controller.AddModuleToCourse(uint(courseID), &moduleRequest)
	if err != nil {
		http.Error(w, "Error adding module to course", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CourseHandler) AddSubjectToModule(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	if requestContext.Role != models.TeacherRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	moduleIDString := r.PathValue("id")
	moduleID, _ := strconv.ParseUint(moduleIDString, 10, 64)

	var subjectRequest controllers.SubjectRequest
	err := json.NewDecoder(r.Body).Decode(&subjectRequest)
	if err != nil {
		http.Error(w, "Error decoding subject request", http.StatusBadRequest)
		return
	}
	err = h.controller.AddSubjectToModule(uint(moduleID), &subjectRequest)
	if err != nil {
		http.Error(w, "Error adding subject to module", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CourseHandler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	courseIDString := r.PathValue("id")
	courseID, _ := strconv.ParseUint(courseIDString, 10, 64)
	course, err := h.controller.GetCourseByID(uint(courseID))
	if err != nil {
		http.Error(w, "Error getting course", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(course)
	if err != nil {
		http.Error(w, "Error encoding course", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CourseHandler) GetAllCoursesByStudentID(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context().Value(RequestContextKey).(RequestContext)
	if requestContext.Role != models.StudentRole {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	studentID := requestContext.UserID
	courses, err := h.controller.GetAllCoursesByStudentID(studentID)
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
