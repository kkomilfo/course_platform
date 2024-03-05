package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Administrator struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Teacher struct {
	gorm.Model
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	FullName    string    `gorm:"not null"`
	AvatarURL   string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	DateOfBirth time.Time `gorm:"not null"`
	Education   string    `gorm:"not null"`
}

type Student struct {
	gorm.Model
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	AvatarURL string
	FullName  string   `gorm:"not null"`
	Courses   []Course `gorm:"many2many:course_enrollments;"`
}

type Course struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	TeacherID   uint      `gorm:"not null"`
	Teacher     Teacher   `gorm:"not null"`
	Modules     []Module  `gorm:"foreignKey:CourseID"`
	Students    []Student `gorm:"many2many:course_enrollments;"`
}

type SubjectType string

// Declare the valid constants for SubjectType
const (
	Lecture SubjectType = "Lecture"
	Task    SubjectType = "Task"
)

type Subject struct {
	gorm.Model
	Title        string `gorm:"not null"`
	Description  string
	Files        []File `gorm:"many2many:subject_files;"`
	StudentFiles []File `gorm:"many2many:subject_student_files;"`
	DueDate      time.Time
	Type         SubjectType
	ModuleID     uint
}

type File struct {
	gorm.Model
	URL string
}

type Module struct {
	gorm.Model
	Title    string
	CourseID uint
	Subjects []Subject `gorm:"foreignKey:ModuleID"`
}

type Grade struct {
	gorm.Model
	StudentID uint
	SubjectID uint
	Grade     float32
}

type Comment struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	UserType  string `gorm:"type:varchar(10);not null"` // 'student' or 'teacher'
	SubjectID uint   `gorm:"not null"`
	Content   string `gorm:"not null"`
	Subject   Subject
}

func main() {
	dsn := fmt.Sprintf("host=localhost user=postgres password=QazWsx@Edc1234 dbname=course_platform port=5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("Started good")

	err = db.AutoMigrate(
		&Teacher{},
		&Administrator{},
		&Student{},
		&Course{},
		&Subject{},
		&Grade{},
		&Comment{},
		&File{},
		&Module{},
	)
	if err != nil {
		return
	}

	// Create a product
}
