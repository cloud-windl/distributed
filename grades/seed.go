package grades

import "gorm.io/gorm"

func SeedStudents(db *gorm.DB) error {
	students := []StudentModel{
		{FirstName: "Tom", LastName: "Jerry"},
		{FirstName: "Alice", LastName: "Smith"},
		{FirstName: "Bob1", LastName: "Johnson"},
		{FirstName: "Bob2", LastName: "Johnson"},
		{FirstName: "Bob3", LastName: "Johnson"},
		{FirstName: "Bob4", LastName: "Johnson"},
		{FirstName: "Bob5", LastName: "Johnson"},
		{FirstName: "Bob6", LastName: "Johnson"},
		{FirstName: "Bob7", LastName: "Johnson"},
		{FirstName: "Bob8", LastName: "Johnson"},
		{FirstName: "Bob9", LastName: "Johnson"},
		{FirstName: "Bob10", LastName: "Johnson"},
	}

	for _, s := range students {
		var count int64
		if err := db.Model(&StudentModel{}).
			Where("first_name = ? AND last_name = ?", s.FirstName, s.LastName).
			Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			if err := db.Create(&s).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
