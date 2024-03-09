package models

type Module struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	Title    string    `gorm:"not null" json:"title"`
	CourseID uint      `gorm:"not null" json:"course_id"`
	Subjects []Subject `gorm:"foreignKey:ModuleID" json:"subjects"`
}
