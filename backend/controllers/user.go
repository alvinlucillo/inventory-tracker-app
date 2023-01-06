package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"personal-inventory/backend/models"
	"personal-inventory/backend/repo"
	utils "personal-inventory/backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s UserService) Login(c *gin.Context) {
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
