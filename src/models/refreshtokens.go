package models

type RefreshToken struct {
	Base
	UserID uint
	Token  string
	User   User
}
