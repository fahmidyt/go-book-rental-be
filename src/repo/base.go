package repo

import (
	"errors"
	"reflect"

	"github.com/fahmidyt/go-book-rental-be/src/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepo struct {
	Model interface{}
}

func (repo BaseRepo) GetAll() (interface{}, error) {
	dType := reflect.TypeOf(repo.Model)
	makeSlice := reflect.New(reflect.SliceOf(dType)).Interface()

	res := db.GetDB().Model(&repo.Model).Preload(clause.Associations).Find(makeSlice)

	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return makeSlice, res.Error
	}

	return makeSlice, nil
}

func (repo BaseRepo) GetOne(id uint) (interface{}, error) {
	res := db.GetDB().First(&repo.Model)

	if res.Error != nil {
		return &repo.Model, res.Error
	}

	return &repo.Model, res.Error
}

func (repo BaseRepo) Create(payload interface{}) (data interface{}, err error) {
	res := db.GetDB().Create(&payload)

	if res.Error != nil {
		return payload, res.Error
	}

	return payload, err
}

func (repo BaseRepo) Update(id uint, payload interface{}) (data interface{}, err error) {
	res := db.GetDB().First(&repo.Model, id)

	if res.Error != nil {
		return data, res.Error
	}

	res.Updates(payload)

	if res.Error != nil {
		return &repo.Model, res.Error
	}

	return &repo.Model, nil
}
