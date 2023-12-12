package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Enabled  bool `json:"enabled"`
	Bank     Bank `gorm:"foreignkey:UserId;association_foreignkey:ID;" json:"bank"`
}
