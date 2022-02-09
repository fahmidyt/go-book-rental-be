package models

type Role struct {
	Base
	Name string `gorm:"unique"`
}
