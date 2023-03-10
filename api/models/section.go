package models

type Section struct {
	Id      int    `json:"id" gorm:"PrimaryKey"`
	Color   string `json:"color" validate:"required"`
	Label   string `json:"label" validate:"required"`
	Xmini   int    `json:"xmini" validate:"required"`
	Xmaxi   int    `json:"xmaxi" validate:"required"`
	Ymini   int    `json:"ymini" validate:"required"`
	Ymaxi   int    `json:"ymaxi" validate:"required"`
	ImageId int    `json:"image_id" validate:"required"`
}

type SectionRequest struct {
}
