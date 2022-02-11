package repo

import (
	"errors"

	"github.com/fahmidyt/go-book-rental-be/src/db"
	"gorm.io/gorm"
)

// TODO: reflect interface{} and make it slice
type BaseRepo struct {
	Model interface{}
}

func (repo BaseRepo) GetAll() ([]map[string]interface{}, error) {
	mockModel := []map[string]interface{}{}
	res := db.GetDB().Model(&repo.Model).Find(&mockModel)

	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return mockModel, res.Error
	}

	return mockModel, nil
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
