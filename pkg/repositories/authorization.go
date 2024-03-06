package repositories

import (
	"awesomeProject/pkg/models"
	"gorm.io/gorm"
)

type AuthorizationRepository struct {
	db *gorm.DB
}

func NewAuthorizationRepository(db *gorm.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db}
}

func (r *AuthorizationRepository) FindAdministratorByEmail(email string) (models.Administrator, error) {
	var user models.Administrator
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *AuthorizationRepository) FindTeacherByEmail(email string) (models.Teacher, error) {
	var user models.Teacher
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *AuthorizationRepository) FindStudentByEmail(email string) (models.Student, error) {
	var user models.Student
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
