package utils

import (
	"errors"

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

type retryFn func() (tryAgain bool)

func Retry(fn retryFn, maxRetries int) error {

	var tryAgain bool
	tries := 1

	for {
		tryAgain = fn()
		if !tryAgain {
			return nil
		}

		tries++
		if tries > maxRetries {
			return errors.New("maximum retries have been reached")
		}
	}
}
