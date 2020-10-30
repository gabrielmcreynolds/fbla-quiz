package usecases

import (
	"backend/errorCodes"
	entity2 "backend/questions/entity"
)

type Service interface {
	GetQuestions() ([]*entity2.Question, *errorCodes.Slug)
}
