package handlers

import (
	"awesomeProject/pkg/controllers"
	"awesomeProject/pkg/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

type userType string

const (
	userTypeTeacher       userType = "teacher"
	userTypeAdministrator userType = "administrator"
	userTypeStudent       userType = "student"
)

const RequestContextKey string = "RequestContextKey"

type RequestContext struct {
	UserID uint
	Role   models.Role
}

type AuthorizationHandler struct {
	controller *controllers.AuthorizationController
}

func NewAuthorizationHandler(controller *controllers.AuthorizationController) *AuthorizationHandler {
	return &AuthorizationHandler{controller}
}

func (h *AuthorizationHandler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		fmt.Println(tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userID := uint(claims["user_id"].(float64))
		role := claims["role"].(string)

		requestContext := RequestContext{UserID: userID, Role: models.Role(role)}
		ctx := context.WithValue(r.Context(), RequestContextKey, requestContext)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
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
