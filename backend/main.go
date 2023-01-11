package main

import (
	"log"
	"os"
	"personal-inventory/backend/controllers"
	"personal-inventory/backend/models"
	"personal-inventory/backend/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	db := initDB()

	inventoryService := controllers.InventoryService{
		DB:          db,
		ServicePort: port,
		Engine:      engine,
	}

	if err := inventoryService.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDB() *gorm.DB {

	var db *gorm.DB
	var err error

	maxRetries := 5
	err = utils.Retry(func() bool {
		time.Sleep(1 * time.Second)

		db, err = gorm.Open(postgres.Open("host=database user=postgres password=postgres dbname=inventorydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
		if err != nil {
			log.Println(err)
			return true
		}

		return false
	}, maxRetries)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(models.User{}); err != nil {
		log.Fatal(err)
	}

	return db
}
