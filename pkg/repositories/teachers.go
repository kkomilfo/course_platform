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
	err := r.db.Preload("Courses").Find(&teachers).Error
	return teachers, err
}

func (r *TeacherRepository) GradeStudentWork(studentWorkID uint, grade int) error {
	return r.db.
		Model(&models.StudentWork{ID: studentWorkID}).
		Update("grade", grade).
		Error
}
