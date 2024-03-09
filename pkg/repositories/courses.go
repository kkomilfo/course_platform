package repositories

import (
	"awesomeProject/pkg/models"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db}
}

func (r *CourseRepository) Create(course *models.Course) error {
	return r.db.Create(course).Error
}
