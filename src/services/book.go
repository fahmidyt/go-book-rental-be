package services

import (
	"github.com/fahmidyt/go-book-rental-be/src/models"
	"github.com/fahmidyt/go-book-rental-be/src/repo"
)

type BookService struct {
	repo.BaseRepo
}

var serviceModel = new(models.Book)

func InitService() *BookService {
	p := new(BookService)
	p.Model = serviceModel

	return p
}
