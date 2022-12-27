package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/koybigino/food-app/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatalf("Some error occured. Err: %s", err)
	//}

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", os.Getenv("DATABASE_URL"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection Error")
	}

	fmt.Println("Database Connection succed !")

	if err := db.AutoMigrate(&models.User{}, &models.Image{}, &models.Section{}, &models.Date{}, &models.Disease{}, &models.Food{}); err != nil {
		fmt.Println(err.Error())
		panic("Error ro create the table")
	}

	return db
}
