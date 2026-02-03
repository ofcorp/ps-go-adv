package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone    string `gorm:"indexUnique"`
	SessionID string `gorm:"indexUnique"`
	VerificationCode string
}
