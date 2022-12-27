package models

type FoodRequest struct {
	Name                string  `json:"name" validate:"required"`
	Quantity            float32 `json:"quantity" validate:"required"`
	Number              int     `json:"number" validate:"required"`
	QuantityWater       float32 `json:"quantity_water" validate:"required"`
	QuantityOtherLiquid float32 `json:"quantity_other_liquid" validate:"required"`
	SportDuration       int     `json:"sport_duration" validate:"required"`
	FruitLegume         bool    `json:"fruit_Legume"  validate:"required"`
	DiseaseName         string  `json:"disease_name" validate:"required" gorm:"unique"`
	Year                int     `json:"year" validate:"required"`
	Month               int     `json:"month" validate:"required"`
	Day                 int     `json:"day" validate:"required"`
	UserId              int     `json:"user_id"`
}

type Food struct {
	Id                  int       `json:"id" gorm:"PrimaryKey"`
	Name                string    `json:"name"`
	Quantity            float32   `json:"quantity"`
	Number              int       `json:"number"`
	QuantityWater       float32   `json:"quantity_water"`
	QuantityOtherLiquid float32   `json:"quantity_other_liquid"`
	SportDuration       int       `json:"sport_duration"`
	FruitLegume         bool      `json:"fruit_Legume" `
	UserId              int       `json:"user_id"`
	Diseases            []Disease `json:"diseases" gorm:"many2many:food_diseases"`
	Dates               []Date    `json:"dates" gorm:"many2many:food_dates"`
}

func ParseFoodRequestToFood(fr FoodRequest, f *Food) {
	f.Name = fr.Name
	f.Quantity = fr.Quantity
	f.Number = fr.Number
	f.QuantityWater = fr.QuantityWater
	f.QuantityOtherLiquid = fr.QuantityOtherLiquid
	f.SportDuration = fr.SportDuration
	f.FruitLegume = fr.FruitLegume
	f.UserId = fr.UserId
}

func ParseFoodRequestToDate(fr FoodRequest, d *DateRequest) {
	d.Year = fr.Year
	d.Month = fr.Month
	d.Day = fr.Day
}

func ParseFoodRequestToDisease(fr FoodRequest, d *Disease) {
	d.Name = fr.DiseaseName
}
