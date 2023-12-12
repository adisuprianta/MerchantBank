package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	UserId uint  `gorm:"unique" json:"user_id"`
	Amount int64 `json:"amount"`
}
