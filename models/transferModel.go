package models

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	UserId      uint   `json:"user_id"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}
