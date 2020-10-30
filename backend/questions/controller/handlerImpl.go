package controller

import (
	"backend/helpers"
	"backend/questions/usecases"
	"github.com/labstack/echo"
	"net/http"
)

type controller struct {
	questionService usecases.Service
}

func (con controller) GetFiveQuestions() func(c echo.Context) error {
	return func(c echo.Context) error {
		questions, err := con.questionService.GetQuestions()
		if err != nil {
			return err.Response(&c)
		}
		return c.JSON(http.StatusOK, helpers.Json{
			"questions": questions,
		})
	}
}

func NewQuestionController(service usecases.Service) Handler {
	return &controller{
		questionService: service,
	}
}
