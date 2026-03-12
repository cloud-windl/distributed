package grades

import "gorm.io/gorm"

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

	if count == 0 {
		users := []UserModel{
			{Username: "admin", Password: "123456", Role: "admin"},
			{Username: "teacher", Password: "123456", Role: "teacher"},
		}
		if err := db.Create(&users).Error; err != nil {
			return err
		}
	}

	return nil
}
