package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Enabled  bool
	Transfer []Transfer `gorm:"foreignkey:UserId;association_foreignkey:ID;" json:"transfer"`
}
