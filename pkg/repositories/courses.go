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

func (r *CourseRepository) FindCourseByID(courseID uint) (models.Course, error) {
	var course models.Course
	err := r.db.Preload("Modules.Subjects.Files").First(&course, courseID).Error
	return course, err
}

func (r *CourseRepository) FindAllByStudentID(id uint) ([]models.Course, error) {
	var courses []models.Course
	err := r.db.
		Joins("JOIN course_enrollments ON course_enrollments.course_id = courses.id AND course_enrollments.student_id = ?", id).
		Find(&courses).
		Error
	return courses, err
}

func (r *CourseRepository) FindSubject(subjectID uint) (models.Subject, error) {
	var subject models.Subject
	err := r.db.
		Preload("Files").
		First(&subject, subjectID).
		Error
	return subject, err
}

func (r *CourseRepository) FindStudentWork(subjectID uint, studentID uint) (models.StudentWork, error) {
	var work models.StudentWork
	err := r.db.
		Preload("Files").
		Preload("Comments").
		Where("subject_id = ? AND student_id = ?", subjectID, studentID).
		Find(&work).
		Error
	return work, err
}

func (r *CourseRepository) FindAllEntrolledStudentsByCourseID(id uint) ([]models.Student, error) {
	var students []models.Student
	err := r.db.
		Joins("JOIN course_enrollments ON course_enrollments.student_id = students.id AND course_enrollments.course_id = ?", id).
		Find(&students).
		Error
	return students, err
}

func (r *CourseRepository) GetStudentsWorksWithSubject(studentID uint, subjectIDs []uint) ([]models.StudentWork, error) {
	var studentWorks []models.StudentWork
	result := r.db.Where("student_id = ? AND subject_id IN ?", studentID, subjectIDs).Find(&studentWorks)
	if result.Error != nil {
		return nil, result.Error
	}
	return studentWorks, nil
}
