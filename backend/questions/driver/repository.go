package driver

import (
	"backend/errorCodes"
	"backend/questions/entity"
)

// Driver for database actions
type Repository interface {

	// Gets five random question from the database
	GetFiveQuestions() ([]*entity.Question, *errorCodes.Slug)
}
