package models

type User struct {
	Base
	Email      string `gorm:"unique;size:100"`
	Password   string `gorm:"not null;size:100"`
	RoleId     uint
	Active     bool
	Role       Role
	UserDetail UserDetail
}
