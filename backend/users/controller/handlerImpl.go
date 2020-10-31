package controller

import (
	"backend/helpers"
	"backend/users/entity"
	"backend/users/usecases"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type controller struct {
	userService usecases.Service
}

func NewUserController(service usecases.Service) Handler {
	return &controller{
		userService: service,
	}
}

func (con controller) CreateUser() func(c echo.Context) error {
	return func(c echo.Context) error {
		auth := new(entity.Authentication)
		if err := c.Bind(auth); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseError{
				Message:    "Invalid Request",
				Resolution: "Make sure the body is JSON with email, name, and password fields",
				Error:      err,
			})
		}

		user, err := con.userService.CreateUser(auth)
		if err != nil {
			return err.Response(&c)
		}

		refreshToken, accessToken, err := con.userService.AddRefreshToken(user, c.RealIP())
		if err != nil {
			return err.Response(&c)
		}

		refreshTokenString, err := refreshToken.GenerateJWT()
		if err != nil {
			return err.Response(&c)
		}
		accessTokenString, err := accessToken.GenerateJWT()
		if err != nil {
			return err.Response(&c)
		}

		return c.JSON(http.StatusCreated, helpers.Json{
			"user":         user,
			"refreshToken": refreshTokenString,
			"accessToken":  accessTokenString,
		})
	}
}

func (con controller) Login() func(c echo.Context) error {
	return func(c echo.Context) error {
		auth := new(entity.Authentication)
		if err := c.Bind(auth); err != nil || c.Validate(auth) != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseError{
				Message:    "Invalid Request",
				Resolution: "Make sure the body is JSON with email and password fields",
				Error:      err,
			})
		}

		user, err := con.userService.Validate(auth)
		if err != nil {
			return err.Response(&c)
		}

		refreshToken, accessToken, err := con.userService.AddRefreshToken(user, c.RealIP())
		if err != nil {
			return err.Response(&c)
		}

		refreshTokenString, err := refreshToken.GenerateJWT()
		if err != nil {
			return err.Response(&c)
		}
		accessTokenString, err := accessToken.GenerateJWT()
		if err != nil {
			return err.Response(&c)
		}

		return c.JSON(http.StatusCreated, helpers.Json{
			"user":         user,
			"refreshToken": refreshTokenString,
			"accessToken":  accessTokenString,
		})
	}
}

func (con controller) Refresh() func(c echo.Context) error {
	return func(c echo.Context) error {
		body := new(struct {
			RefreshToken string `json:"refreshToken"`
		})
		if err := c.Bind(body); err != nil || c.Validate(body) != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseError{
				Message:    "Must contain refreshToken in body",
				Resolution: "Include 'refreshToken: \"your_refresh_token\"' in your body as JSON",
			})
		}
		accessToken, err := con.userService.RefreshToken(body.RefreshToken)
		if err != nil {
			return err.Response(&c)
		}

		accessTokenString, err := accessToken.GenerateJWT()
		if err != nil {
			return err.Response(&c)
		}
		return c.JSON(http.StatusCreated, helpers.Json{
			"message":     "Created Successfully!",
			"accessToken": accessTokenString,
		})
	}
}

func (con controller) Logout() func(c echo.Context) error {
	return func(c echo.Context) error {
		body := new(struct {
			RefreshToken string `json:"refreshToken"`
		})
		if err := c.Bind(body); err != nil || c.Validate(body) != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseError{
				Message:    "Must contain refreshToken in body",
				Resolution: "Include 'refreshToken: \"your_refresh_token\"' in your body as JSON",
			})
		}

		err := con.userService.RemoveRefreshToken(body.RefreshToken)
		if err != nil {
			return err.Response(&c)
		}
		return c.JSON(http.StatusOK, helpers.Json{
			"message": "successfully logged out user",
		})
	}
}

func (con controller) AddTest() func(c echo.Context) error {
	return func(c echo.Context) error {
		body := new(entity.TestResult)
		if err := c.Bind(body); err != nil || c.Validate(body) != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseError{
				Message: "Must be a contain a score and duration in body",
			})
		}
		idString := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["userId"]
		id, _ := primitive.ObjectIDFromHex(idString.(string))
		fmt.Printf("id: %v", id)
		user, err := con.userService.AddTest(&id, body)
		if err != nil {
			return err.Response(&c)
		}
		return c.JSON(http.StatusCreated, helpers.Json{
			"message": "successfully added test",
			"user":    user,
		})
	}
}
