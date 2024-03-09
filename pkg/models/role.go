package models

type Role string

const (
	TeacherRole       Role = "teacher"
	AdministratorRole Role = "admin"
	StudentRole       Role = "student"
)
