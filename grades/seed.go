package grades

import "gorm.io/gorm"

func SeedStudents(db *gorm.DB) error {
	var count int64
	if err := db.Model(&StudentModel{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	students := []StudentModel{
		{FirstName: "Tom", LastName: "Jerry"},
		{FirstName: "Alice", LastName: "Smith"},
		{FirstName: "Bob", LastName: "Johnson"},
	}

	return db.Create(&students).Error
}
