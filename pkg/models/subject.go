package models

import (
	"time"
)

type SubjectType string

const (
	Lecture SubjectType = "Lecture"
	Task    SubjectType = "Task"
)

type Subject struct {
	ID           uint        `gorm:"primarykey" json:"id"`
	Title        string      `gorm:"not null" json:"title"`
	Description  string      ` json:"description"`
	Files        []File      `gorm:"many2many:subject_files;" json:"files"`
	StudentFiles []File      `gorm:"many2many:subject_student_files;" json:"student_files"`
	DueDate      time.Time   `json:"due_date"`
	Type         SubjectType `json:"type"`
	ModuleID     uint        `json:"module_id"`
}

type File struct {
	ID  uint   `gorm:"primarykey" json:"id"`
	URL string `gorm:"not null" json:"url"`
}
