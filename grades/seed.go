package grades

import (
	"distributed/pkg/password"

	"gorm.io/gorm"
)

func SeedStudents(db *gorm.DB) error {
	var count int64
	if err := db.Model(&StudentModel{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		students := []StudentModel{
			{FirstName: "Tom", LastName: "Jerry"},
			{FirstName: "Alice", LastName: "Smith"},
			{FirstName: "Bob", LastName: "Johnson"},
		}
		if err := db.Create(&students).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedUsers(db *gorm.DB) error {
	var count int64
	if err := db.Model(&UserModel{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	adminPassword, err := password.HashPassword("123456")
	if err != nil {
		return err
	}

	teacherPassword, err := password.HashPassword("123456")
	if err != nil {
		return err
	}

	users := []UserModel{
		{Username: "admin", Password: adminPassword, Role: "admin"},
		{Username: "teacher", Password: teacherPassword, Role: "teacher"},
	}

	return db.Create(&users).Error
}
