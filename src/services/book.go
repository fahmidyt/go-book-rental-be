package services

import (
	"github.com/fahmidyt/go-book-rental-be/src/models"
	"github.com/fahmidyt/go-book-rental-be/src/repo"
)

type BookService struct {
	repo.BaseRepo
}

var bookModel = new(models.Book)

func InitBookService() *BookService {
	p := new(BookService)
	p.Model = bookModel

	return p
}
