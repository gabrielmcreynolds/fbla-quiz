package controller

import "github.com/labstack/echo"

type Handler interface {
	GetFiveQuestions() func(c echo.Context) error
}
