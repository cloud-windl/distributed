package grades

import "time"

type StudentModel struct {
	ID        int64        `gorm:"primaryKey;autoIncrement"`
	FirstName string       `gorm:"type:varchar(64);not null"`
	LastName  string       `gorm:"type:varchar(64);not null"`
	Grades    []GradeModel `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GradeModel struct {
	ID        int64   `gorm:"primaryKey;autoIncrement"`
	StudentID int64   `gorm:"index;not null"`
	Title     string  `gorm:"type:varchar(128);not null"`
	Type      string  `gorm:"type:varchar(32);not null"`
	Score     float32 `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
