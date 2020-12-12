package controller

import (
	"github.com/labstack/echo"
)

// Handler interfaces http request to functions from [usecases] package
// All methods are http calls
type Handler interface {
	CreateUser() func(c echo.Context) error
	GetUser() func(c echo.Context) error
	Login() func(c echo.Context) error
	Refresh() func(c echo.Context) error
	Logout() func(c echo.Context) error
	AddTest() func(c echo.Context) error
}
