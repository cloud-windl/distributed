package grades

import "time"

type UserModel struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"type:varchar(64);uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Role      string `gorm:"type:varchar(32);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
