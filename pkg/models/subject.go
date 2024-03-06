package models

import (
	"gorm.io/gorm"
	"time"
)

type SubjectType string

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
