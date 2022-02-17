package models

type User struct {
	Base
	Email      string `gorm:"unique;size:100" json:"email"`
	Password   string `gorm:"not null;size:100" json:"password"`
	RoleId     uint   `json:"RoleId"`
	Active     bool   `json:"active"`
	Role       Role
	UserDetail UserDetail
}
