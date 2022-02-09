package models

type RentedBookDetail struct {
	Base
	RentedBookID uint
	BookID       uint
	Book         Book
}
