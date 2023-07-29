package models

import "gorm.io/gorm"

type Animal struct {
	gorm.Model
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Nickname string `json:"nickname"`
}
