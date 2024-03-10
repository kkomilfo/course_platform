package db

import (
	"awesomeProject/pkg/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDatabase() *gorm.DB {
	host := os.Getenv("DATABASE_HOST")
	password := os.Getenv("DATABASE_PASSWORD")
	userName := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, userName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database")
	err = db.AutoMigrate(
		&models.Teacher{},
		&models.Student{},
		&models.Course{},
		&models.Subject{},
		&models.Grade{},
		&models.StudentWork{},
		&models.Comment{},
		&models.StudentWorkFile{},
		&models.File{},
		&models.Module{},
		&models.Administrator{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}
	return db
}
