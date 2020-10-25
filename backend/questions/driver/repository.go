package driver

import (
	"backend/errorCodes"
	"backend/questions/entity"
)

type Repository interface {

	// Gets five random question from the database
	GetFiveQuestions() ([]*entity.Question, *errorCodes.Slug)
}
