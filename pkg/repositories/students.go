package repositories

import (
	"awesomeProject/pkg/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db}
}

func (r *StudentRepository) Create(student *models.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepository) FindAll() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Find(&students).Error
	return students, err
}

func (r *StudentRepository) CreateStudentWork(work *models.StudentWork) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.
			Create(work).
			Error

		if err != nil {
			return err
		}

		for i := range work.Files {
			work.Files[i].StudentWorkID = work.ID
		}

		return nil
	})
}

func (r *StudentRepository) Comment(comment *models.Comment) error {
	return r.db.Create(comment).Error
}
