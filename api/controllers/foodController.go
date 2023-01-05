package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
	"github.com/koybigino/food-app/api/validations"
)

func GetAllFoods(c *fiber.Ctx) error {
	var Foods []models.Food

	if err := db.Preload("FoodEats").Find(&Foods).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Foods": Foods,
	})
}

func GetAllFoodNumber(c *fiber.Ctx) error {
	var Foods []models.Food
	var number int

	if err := db.Preload("FoodEats").Find(&Foods).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	number = 0

	for _, f := range Foods {
		for _, val := range f.FoodEats {
			number = number + val.Number
		}
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"number": number,
	})

}

func GetTheFoodNumber(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	body := new(models.Food)
	var number int

	if err := db.Preload("FoodEats").First(&body, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	number = len(body.FoodEats)

	for _, val := range body.FoodEats {
		number = number + (val.Number - 1)
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"number": number,
	})

}

func GetFoodById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	body := new(models.Food)

	if err := db.Preload("FoodEats").First(&body, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	for i := 0; i < len(body.FoodEats); i++ {
		fe := models.FoodEat{}
		fe.Id = body.FoodEats[i].Id
		if err := db.Preload("Dates").Preload("Diseases").First(&fe).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad Credentials !",
			})
		}

		body.FoodEats[i].Dates = fe.Dates
		body.FoodEats[i].Diseases = fe.Diseases

	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Food": body,
	})
}

func CreateFood(c *fiber.Ctx) error {
	body := new(models.FoodRequest)
	food := new(models.Food)
	foodEat := new(models.FoodEat)
	dateRequest := new(models.DateRequest)
	date := new(models.Date)
	disease := new(models.Disease)
	foodDate := new(models.FoodDates)
	foodDisease := new(models.FoodDiseases)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	errors := validations.ValidateStruct(body)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	models.ParseFoodRequestToDate(*body, dateRequest)
	models.ParseFoodRequestToFood(*body, food)
	models.ParseFoodRequestToDisease(*body, disease)
	models.GenerateDate(*dateRequest, date)
	models.ParseFoodRequestToFoodEat(*body, foodEat)

	if err := db.Create(food).Error; err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"foods_name_key\" (SQLSTATE 23505)" {
			db.Where("name = ?", food.Name).First(food)
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	if food.Id == 0 {
		food = &models.Food{}

		db.Last(food)
	}

	foodEat.FoodId = food.Id
	if err := db.Create(foodEat).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	if err := db.Create(date).Error; err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"dates_date_key\" (SQLSTATE 23505)" {
			db.Where("date = ?", date.Date).First(date)
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	if err := db.Create(disease).Error; err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"diseases_name_key\" (SQLSTATE 23505)" {
			db.Where("name = ?", disease.Name).First(disease)
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	if date.Id == 0 {
		date = &models.Date{}

		db.Last(date)
	}

	if disease.Id == 0 {
		disease = &models.Disease{}

		db.Last(disease)
	}

	foodEat = &models.FoodEat{}

	db.Last(foodEat)

	foodDate.DateId = date.Id
	foodDate.FoodEatId = foodEat.Id
	if err := db.Create(foodDate).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	foodDisease.DiseaseId = disease.Id
	foodDisease.FoodEatId = foodEat.Id
	if err := db.Create(foodDisease).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Food creation ok !",
		"Food":    body,
	})
}

func DeleteFood(c *fiber.Ctx) error {
	section := new(models.Food)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := db.Where("image_id = ?", id).Delete(&models.FoodDates{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	if err := db.Where("image_id = ?", id).Delete(&models.FoodDiseases{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	if err := db.Delete(&section, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
