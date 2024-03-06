package controllers

import (
	"awesomeProject/pkg/repositories"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var ErrInternal = errors.New("internal error")
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthorizationController struct {
	repository *repositories.AuthorizationRepository
}

func NewAuthorizationController(repository *repositories.AuthorizationRepository) *AuthorizationController {
	return &AuthorizationController{repository}
}

func (c *AuthorizationController) LoginAsTeacher(credentials Credentials) (string, error) {
	teacher, err := c.repository.FindTeacherByEmail(credentials.Email)
	if err != nil {
		return "", ErrUserNotFound
	}
	if credentials.Password != teacher.Password {
		return "", ErrInvalidCredentials
	}
	tokenString, err := c.generateToken(teacher.ID)
	if err != nil {
		return "", ErrInternal
	}
	return tokenString, nil
}

func (c *AuthorizationController) LoginAsAdministrator(credentials Credentials) (string, error) {
	admin, err := c.repository.FindAdministratorByEmail(credentials.Email)
	if err != nil {
		return "", ErrUserNotFound
	}
	if credentials.Password != admin.Password {
		return "", ErrInvalidCredentials
	}
	tokenString, err := c.generateToken(admin.ID)
	if err != nil {
		return "", ErrInternal
	}
	return tokenString, nil
}

func (c *AuthorizationController) LoginAsStudent(credentials Credentials) (string, error) {
	student, err := c.repository.FindStudentByEmail(credentials.Email)
	if err != nil {
		return "", ErrUserNotFound
	}
	if credentials.Password != student.Password {
		return "", ErrInvalidCredentials
	}
	tokenString, err := c.generateToken(student.ID)
	if err != nil {
		return "", ErrInternal
	}
	return tokenString, nil
}

func (c *AuthorizationController) generateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", ErrInternal
	}

	return tokenString, nil
}
