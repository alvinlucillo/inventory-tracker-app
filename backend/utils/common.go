package utils

import (
	"gorm.io/gorm"
)

type DBError struct {
	Error error
	Model string
}

func HandleDBError(db *gorm.DB) *DBError {
	if db.Error != nil {
		return &DBError{
			Error: db.Error,
			Model: db.Statement.Table,
		}
	}

	return nil
}
