package controller

import "github.com/labstack/echo"

// Handler interfaces http to usecases
type Handler interface {
	// Gets Five randomly unique questions
	GetFiveQuestions() func(c echo.Context) error
}
