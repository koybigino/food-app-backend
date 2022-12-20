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

	port, _ := strconv.Atoi(os.Getenv("food-app-env-var/Port"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", os.Getenv("food-app-env-var/Hostname"), os.Getenv("food-app-env-var/Username"), os.Getenv("food-app-env-var/Password"), os.Getenv("food-app-env-var/Database"), port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection Error")
	}

	fmt.Println("Connection succed !")

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Error ro create the table")
	}

	return db
}
