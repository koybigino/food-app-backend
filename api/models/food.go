package models

type FoodRequest struct {
	Name                string  `json:"name" validate:"required"`
	Quantity            float32 `json:"quantity" validate:"required"`
	NutritiveValue      float32 `json:"nutritive_value" validate:"required"`
	Number              int     `json:"number" validate:"required"`
	QuantityWater       float32 `json:"quantity_water" validate:"required"`
	QuantityOtherLiquid float32 `json:"quantity_other_liquid" validate:"required"`
	SportDuration       int     `json:"sport_duration" validate:"required"`
	FruitLegume         bool    `json:"fruit_Legume"  validate:"required"`
	DiseaseName         string  `json:"disease_name" validate:"required" gorm:"unique"`
	Year                int     `json:"year" validate:"required"`
	Month               int     `json:"month" validate:"required"`
	Day                 int     `json:"day" validate:"required"`
	UserId              int     `json:"user_id" validate:"required"`
}

type FoodEat struct {
	Id                  int       `json:"id" gorm:"PrimaryKey"`
	Number              int       `json:"number"`
	QuantityWater       float32   `json:"quantity_water"`
	QuantityOtherLiquid float32   `json:"quantity_other_liquid"`
	SportDuration       int       `json:"sport_duration"`
	FruitLegume         bool      `json:"fruit_Legume" `
	FoodId              int       `json:"food_id"`
	Diseases            []Disease `json:"diseases" gorm:"many2many:food_diseases"`
	Dates               []Date    `json:"dates" gorm:"many2many:food_dates"`
}

type Food struct {
	Id             int       `json:"id" gorm:"PrimaryKey"`
	Name           string    `json:"name" gorm:"unique"`
	NutritiveValue float32   `json:"nutritive_value"`
	Quantity       float32   `json:"quantity"`
	UserId         int       `json:"user_id"`
	FoodEats       []FoodEat `json:"food-eat" gorm:"foreignKey:FoodId"`
}

type FoodFill struct {
	Food
	QuantityPerNutritiveValue float32 `json:"quantity_per_nutritive_value"`
}

func ParseFoodRequestToFoodEat(fr FoodRequest, f *FoodEat) {
	f.Number = fr.Number
	f.QuantityWater = fr.QuantityWater
	f.QuantityOtherLiquid = fr.QuantityOtherLiquid
	f.SportDuration = fr.SportDuration
	f.FruitLegume = fr.FruitLegume
}

func ParseFoodRequestToDate(fr FoodRequest, d *DateRequest) {
	d.Year = fr.Year
	d.Month = fr.Month
	d.Day = fr.Day
}

func ParseFoodRequestToDisease(fr FoodRequest, d *Disease) {
	d.Name = fr.DiseaseName
}

func ParseFoodRequestToFood(fr FoodRequest, f *Food) {
	f.Name = fr.Name
	f.UserId = fr.UserId
	f.NutritiveValue = fr.NutritiveValue
	f.Quantity = fr.Quantity
}

func ParseFoodToFoodFill(fr Food, f *FoodFill) {
	f.Quantity = fr.Quantity
	f.Name = fr.Name
	f.UserId = fr.UserId
	f.NutritiveValue = fr.NutritiveValue
	f.QuantityPerNutritiveValue = fr.Quantity / fr.NutritiveValue
}
