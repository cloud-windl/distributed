package grades

import "gorm.io/gorm"

func InitMySQLRepo(db *gorm.DB) (StudentRepo, error) {
	if err := AutoMigrate(db); err != nil {
		return nil, err
	}

	if err := SeedStudents(db); err != nil {
		return nil, err
	}

	if err := SeedUsers(db); err != nil {
		return nil, err
	}

	return NewMySQLStudentRepo(db), nil
}
