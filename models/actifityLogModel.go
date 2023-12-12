package models

type ActivityLog struct {
	Id          uint   `gorm:"primarykey;autoIncrement" json:"id"`
	Action      string `json:"action"`
	Description string `json:"additional_field"`
	UserId      uint
	User        User `gorm:"foreignKey:UserId" json:"user"`
}
