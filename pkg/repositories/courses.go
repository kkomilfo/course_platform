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

func (r *CourseRepository) FindAllByTeacherID(teacherID uint) ([]models.Course, error) {
	var courses []models.Course
	err := r.db.Preload("Students").Where("teacher_id = ?", teacherID).Find(&courses).Error
	return courses, err
}

func (r *CourseRepository) EnrollStudent(studentID uint, courseID uint) error {
	student := models.Student{ID: studentID}
	course := models.Course{ID: courseID}
	return r.db.Model(&student).Association("Courses").Append(&course)
}

func (r *CourseRepository) AddModuleToCourse(courseID uint, module *models.Module) error {
	course := models.Course{ID: courseID}
	return r.db.
		Model(&course).
		Association("Modules").
		Append(module)
}

func (r *CourseRepository) AddSubjectToModule(moduleID uint, subject *models.Subject) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.
			Model(&models.Module{ID: moduleID}).
			Association("Subjects").
			Append(subject)

		if err != nil {
			return err
		}

		for i := range subject.Files {
			subject.Files[i].SubjectID = subject.ID // Assign SubjectID
			if err := tx.Create(&subject.Files[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
