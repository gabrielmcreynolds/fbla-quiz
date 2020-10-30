package usecases

import (
	"backend/errorCodes"
	"backend/questions/driver"
	"backend/questions/entity"
)

type serviceImpl struct {
	repo driver.Repository
}

func (s serviceImpl) GetQuestions() ([]*entity.Question, *errorCodes.Slug) {
	questions, err := s.repo.GetFiveQuestions()
	if err != nil {
		return nil, err
	}
	for checkContainsDup(questions) {
		questions, err = s.repo.GetFiveQuestions()
		if err != nil {
			return nil, err
		}
	}
	return questions, nil
}

// Checks if there is any duplicate questions and if so it will return true
func checkContainsDup(questions []*entity.Question) bool {
	rndmap := make(map[string]bool)
	for i := 0; len(rndmap) < len(questions); i++ {
		rndmap[questions[i].Question] = true
	}
	return len(rndmap) != len(questions)
}

func NewService(repository driver.Repository) Service {
	return &serviceImpl{
		repo: repository,
	}
}
