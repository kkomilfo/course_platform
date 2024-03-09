package repositories

import (
	"awesomeProject/pkg/models"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{db}
}

func (r *TeacherRepository) Create(teacher *models.Teacher) error {
	return r.db.Create(teacher).Error
}

func (r *TeacherRepository) FindAll() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := r.db.Find(&teachers).Error
	return teachers, err
}
