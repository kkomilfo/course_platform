package models

type Course struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	TeacherID   uint      `gorm:"not null" json:"teacher_id"`
	Teacher     Teacher   `gorm:"not null" json:"teacher"`
	Modules     []Module  `gorm:"foreignKey:CourseID" json:"modules"`
	Students    []Student `gorm:"many2many:course_enrollments;" json:"students"`
}
