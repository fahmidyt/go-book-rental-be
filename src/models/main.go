package models

func GetModelStructs() []interface{} {
	var model []interface{}

	model = append(model, &User{})
	model = append(model, &UserDetail{})

	return model
}
