package models

type Disease struct {
	Id   int    `json:"id" gorm:"PrimaryKey"`
	Name string `json:"name" validate:"required" gorm:"unique"`
}
