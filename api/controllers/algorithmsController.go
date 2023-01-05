package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
)

func Predict(c *fiber.Ctx) error {

	day := c.Params("day")
	var dates []models.Date
	totalNum := 0
	foods := []models.Food{}
	foodsC := []models.Food{}
	foodEats := []models.FoodEat{}
	var fes []models.FoodEat
	var fds []models.FoodDates

	if err := db.Find(&foods).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var tableOfProb []float32

	if err := db.Find(&foodEats).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for _, f := range foodEats {
		totalNum = totalNum + f.Number
	}

	if err := db.Where("day = ?", day).Find(&dates).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	fmt.Println(dates)

	for _, d := range dates {
		if err := db.Where("date_id = ?", d.Id).Find(&fds).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad Credentials !",
			})
		}
		fmt.Println(fds)

		for _, f := range fds {
			if err := db.Where("id = ?", f.FoodEatId).Find(&fes).Error; err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "bad Credentials !",
				})
			}
			fmt.Println(fes)

			for _, fe := range fes {
				for _, food := range foods {
					if fe.FoodId == food.Id {
						tableOfProb = append(tableOfProb, (float32(fe.Number))/float32(totalNum))
						foodsC = append(foodsC, food)
					}
				}
			}
		}

	}
	var k int
	k = 0

	for j := 0; j < len(tableOfProb)-1; j++ {
		if tableOfProb[j] > tableOfProb[j+1] {
			k = j
		} else {
			k = j + 1
		}
	}

	if len(foodsC) > 0 {
		return c.JSON(fiber.Map{
			foodsC[k].Name: foodsC[k],
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "vous n'avez pas encore mangé un repas ce jour !",
		})
	}

}

func FillMax(c *fiber.Ctx) error {
	StrQ := c.Params("quantity")
	Q, _ := strconv.Atoi(StrQ)
	var foods []models.Food
	var foodTakes []models.FoodFill
	var foodsFills []models.FoodFill

	if err := db.Preload("FoodEats").Find(&foods).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	for i := 0; i < len(foods); i++ {
		f := models.FoodFill{}
		models.ParseFoodToFoodFill(foods[i], &f)
		foodsFills = append(foodsFills, f)
	}
	fmt.Println(foodsFills)

	for i := 0; i < len(foodsFills)-1; i++ {
		for j := i + 1; j < len(foodsFills); j++ {
			if foodsFills[i].QuantityPerNutritiveValue < foodsFills[j].QuantityPerNutritiveValue {
				food := models.FoodFill{}
				food = foodsFills[i]
				foodsFills[i] = foodsFills[j]
				foodsFills[j] = food
			}

		}
	}
	fmt.Println(foodsFills)

	for _, food := range foodsFills {
		if int(food.Quantity) <= Q {
			foodTakes = append(foodTakes, food)
			Q = Q - int(food.Quantity)
		}
	}

	return c.JSON(fiber.Map{
		"foodsTake": foodTakes,
		"quantity":  Q,
	})
}
