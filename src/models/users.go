package models

type User struct {
	Base
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	RoleId   string `gorm:"not null"`
	Active   bool
}
