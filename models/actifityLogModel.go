package models

import "gorm.io/gorm"

type ActivityLog struct {
	gorm.Model
	Action          string `json:"action"`
	AdditionalField string `json:"additional_field"`
	UserId          uint
	User            User `gorm:"foreignKey:UserId" json:"user"`
}
