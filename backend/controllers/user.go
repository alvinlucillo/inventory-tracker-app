package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"personal-inventory/backend/models"
	"personal-inventory/backend/repo"
	utils "personal-inventory/backend/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InventoryService struct {
	DB          *gorm.DB
	ServicePort string
	Engine      *gin.Engine
}

func (s InventoryService) Run() error {
	s.Engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "OPTIONS", "GET"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization,Access-Control-Allow-Origin"},
	}))

	s.Engine.POST("/auth/login", s.Login)
	s.Engine.POST("/auth/register", s.Register)

	err := s.Engine.Run(":" + s.ServicePort)

	return err
}

func (s InventoryService) Register(c *gin.Context) {
	var userDto models.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRepo := repo.UserRepo{}
	user, err := userRepo.FindByUsername(s.DB, userDto.Username)
	if err != nil {
		// Non-existing username means it can be registered
		if !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, "DB error has occurred. Check DB logs.")
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, "Username already exists.")
		return
	}

	token, jwtErr := utils.GenerateJWT(user.Username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, jwtErr.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}

func (s InventoryService) Login(c *gin.Context) {
	var userDto models.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRepo := repo.UserRepo{}
	user, err := userRepo.FindByUsernamePassword(s.DB, userDto.Username, userDto.Password)
	if err != nil {
		errMsg := "DB error has occurred. Check DB logs."
		statusCode := http.StatusInternalServerError
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			errMsg = "User credentials do not exist."
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, errMsg)

		return
	}

	token, jwtErr := utils.GenerateJWT(user.Username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, jwtErr.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}
