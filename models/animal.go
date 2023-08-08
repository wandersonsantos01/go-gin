package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Animal struct {
	gorm.Model
	Name     string `json:"name" validate:"nonzero"`
	Age      int64  `json:"age" validate:"len=1, regexp=^[0-9]*$"`
	Nickname string `json:"nickname" validate:"nonzero"`
}

func ValidateAnimalData(animal *Animal) error {
	if error := validator.Validate(animal); error != nil {
		return error
	}
	return nil
}
