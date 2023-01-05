package models

import (
	"fmt"
	"strconv"
	"time"
)

type DateRequest struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type Date struct {
	Id   int       `json:"id" gorm:"PrimaryKey"`
	Day  string    `json:"day"`
	Date time.Time `json:"date" gorm:"unique"`
}

func GenerateDate(dr DateRequest, d *Date) {

	var m string
	var day string
	if dr.Month < 10 {
		m = "0" + strconv.Itoa(dr.Month)
	} else {
		m = strconv.Itoa(dr.Month)
	}

	if dr.Day < 10 {
		day = "0" + strconv.Itoa(dr.Month)
	} else {
		day = strconv.Itoa(dr.Month)
	}
	str := fmt.Sprintf("%v-%v-%vT11:45:26.371Z", dr.Year, m, day)

	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		fmt.Println(err)
	}

	d.Date = t
	d.Day = t.Weekday().String()
}
