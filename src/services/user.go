package services

import (
	"github.com/fahmidyt/go-book-rental-be/src/models"
	"github.com/fahmidyt/go-book-rental-be/src/repo"
)

type UserService struct {
	repo.BaseRepo
}

var userModel = new(models.User)

func InitUserService() *UserService {
	p := new(UserService)
	p.Model = userModel

	return p
}
