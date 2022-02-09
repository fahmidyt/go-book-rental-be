package models

import "time"

type UserDetail struct {
	Base
	User        User   `gorm:"embedded"`
	firstName   string `gorm:"not null"`
	lastName    string `gorm:"not null"`
	phoneNumber string `gorm:"not null"`
	dateOfBirth time.Time
}
