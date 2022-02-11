package types

import "github.com/fahmidyt/go-book-rental-be/src/helpers"

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=50"`
}

type RegisterForm struct {
	LoginForm
	ConfirmPassword string       `form:"confirmPassword" json:"confirmPassword" binding:"required,eqfield=Password"`
	FirstName       string       `form:"firstName" json:"firstName" binding:"required,min=3,max=50,min=3,max=50"`
	LastName        string       `form:"lastName" json:"lastName" binding:"required,min=3,max=50"`
	PhoneNumber     string       `form:"phoneNumber" json:"phoneNumber" binding:"required,min=10,max=15"`
	DateOfBirth     helpers.Time `form:"dateOfBirth" json:"dateOfBirth" binding:"required"`
}
