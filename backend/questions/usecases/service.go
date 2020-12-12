package usecases

import (
	"backend/errorCodes"
	entity2 "backend/questions/entity"
)

// Business Logic for Requests
type Service interface {

	// Gets Five Questions, this is where it checks for duplicate questions
	GetQuestions() ([]*entity2.Question, *errorCodes.Slug)
}
