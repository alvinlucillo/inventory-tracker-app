package repo

import (
	"personal-inventory/backend/models"
	"personal-inventory/backend/utils"

	"gorm.io/gorm"
)

type UserRepo struct{}

func (*UserRepo) FindByUsernamePassword(DB *gorm.DB, username string, password string) (*models.User, *utils.DBError) {
	var user *models.User

	if err := utils.HandleDBError(DB.Where("username = ? AND password = ?", username, password).First(&user)); err != nil {
		return nil, err
	}

	return user, nil
}

func (*UserRepo) FindByUsername(DB *gorm.DB, username string) (*models.User, *utils.DBError) {
	var user *models.User

	if err := utils.HandleDBError(DB.Where("username = ?", username).First(&user)); err != nil {
		return nil, err
	}

	return user, nil
}

func (*UserRepo) Create(DB *gorm.DB, user *models.User) *utils.DBError {
	if err := utils.HandleDBError(DB.Create(&user)); err != nil {
		return err
	}

	return nil
}
