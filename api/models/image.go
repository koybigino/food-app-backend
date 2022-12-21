package models

type ImageBase struct {
	Filename string `json:"filename" validate:"required" gorm:"unique"`
	Path     string `json:"path" validate:"required" gorm:"unique"`
}

type Image struct {
	Id int `json:"id" gorm:"PrimaryKey"`
	ImageBase
	Sections []Section `json:"section" gorm:"foreignKey:ImageId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
