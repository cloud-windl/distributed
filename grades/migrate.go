package grades

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&StudentModel{}, &GradeModel{}, &UserModel{})
}
