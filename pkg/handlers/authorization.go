package handlers

import (
	"awesomeProject/pkg/controllers"
	"encoding/json"
	"errors"
	"net/http"
)

type userType string

const (
	userTypeTeacher       userType = "teacher"
	userTypeAdministrator userType = "administrator"
	userTypeStudent       userType = "student"
)

type AuthorizationHandler struct {
	controller *controllers.AuthorizationController
}

func NewAuthorizationHandler(controller *controllers.AuthorizationController) *AuthorizationHandler {
	return &AuthorizationHandler{controller}
}

func (h *AuthorizationHandler) TeacherLogin(w http.ResponseWriter, r *http.Request) {
	h.login(userTypeTeacher, w, r)
}

func (h *AuthorizationHandler) StudentLogin(w http.ResponseWriter, r *http.Request) {
	h.login(userTypeStudent, w, r)
}

func (h *AuthorizationHandler) AdministratorLogin(w http.ResponseWriter, r *http.Request) {
	h.login(userTypeAdministrator, w, r)
}

func (h *AuthorizationHandler) login(u userType, w http.ResponseWriter, r *http.Request) {
	var creds controllers.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Error decoding login request", http.StatusBadRequest)
		return
	}
	var tokenString string
	switch {
	case u == userTypeTeacher:
		tokenString, err = h.controller.LoginAsTeacher(creds)
	case u == userTypeStudent:
		tokenString, err = h.controller.LoginAsStudent(creds)
	case u == userTypeAdministrator:
		tokenString, err = h.controller.LoginAsAdministrator(creds)
	}
	if err != nil {
		if errors.Is(err, controllers.ErrUserNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
		} else if errors.Is(err, controllers.ErrInvalidCredentials) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	if err != nil {
		http.Error(w, "Error encoding login response", http.StatusInternalServerError)
		return
	}
}
