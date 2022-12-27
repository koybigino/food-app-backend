package models

import (
	"fmt"
	"time"
)

type DateRequest struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type Date struct {
	Id   int       `json:"id" gorm:"PrimaryKey"`
	Date time.Time `json:"date" gorm:"unique"`
}

func GenerateDate(dr DateRequest, d *Date) {
	str := fmt.Sprintf("%d-%d-%dT00:00:00.371Z", dr.Year, dr.Month, dr.Day)

	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		fmt.Println(err)
	}

	d.Date = t
}
