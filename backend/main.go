package main

import (
	"log"
	"os"
	"personal-inventory/backend/controllers"
	"personal-inventory/backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	db := initDB()

	userService := controllers.UserService{
		DB: db,
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "OPTIONS", "GET"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization,Access-Control-Allow-Origin"},
	}))

	r.POST("/auth/login", userService.Login)

	r.Run(":" + port)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=inventorydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
