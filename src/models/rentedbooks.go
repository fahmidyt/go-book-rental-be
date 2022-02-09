package models

type RentedBook struct {
	Base
	UserID           uint
	User             User `gorm:"foreignKey:UserID"`
	RentedBookDetail []RentedBookDetail
}
