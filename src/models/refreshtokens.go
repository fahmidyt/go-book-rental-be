package models

type RefreshToken struct {
	Base
	UserID uint
	Token  string `json:"token"`
	User   User
}
