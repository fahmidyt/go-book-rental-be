package types

import "time"

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=50"`
}

type RegisterForm struct {
	FirstName   string    `form:"firstName" json:"firstName" binding:"required,min=3,max=50"`
	LastName    string    `form:"lastName" json:"lastName" binding:"required,min=3,max=50"`
	PhoneNumber string    `form:"phoneNumber" json:"phoneNumber" binding:"required,min=10,max=15"`
	DateOfBirth time.Time `form:"dateOfBirth" json:"dateOfBirth" binding:"required"`
}
