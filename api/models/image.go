package models

type ImageBase struct {
	Filename string `json:"filename" validate:"required"`
	Path     string `json:"path" validate:"required"`
}

type Image struct {
	Id int `json:"id" gorm:"PrimaryKey"`
	ImageBase
	UserId   int       `json:"user_id"`
	Sections []Section `json:"sections" validate:"required" gorm:"foreignKey:ImageId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
