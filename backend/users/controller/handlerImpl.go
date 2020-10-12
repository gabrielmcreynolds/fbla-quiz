package controller

import (
	"backend/errorCodes"
	"backend/helpers"
	"backend/users/entity"
	"backend/users/usecases"
	"github.com/labstack/echo"
	"log"
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

		log.Printf("User%v", auth)

		user, err := con.userService.CreateUser(auth)
		log.Printf("err: %v", err)
		if err != nil {
			switch err.Error() {
			case errorCodes.SameEmailName:
				return c.JSON(http.StatusBadRequest, helpers.ResponseError{
					Message:    "Email already exists in database",
					Resolution: "Use a unique email",
				})
			case errorCodes.PasswordHashingFailure:
				return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
					Message: "Password hashing failure",
					Error:   err,
				})
			default:
				return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
					Message: "could not add user to database",
					Error:   err,
				})
			}
		}

		refreshToken, accessToken, err := con.userService.AddRefreshToken(user, c.RealIP())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
				Message: "Internal error while creating refreshToken",
				Error:   err,
			})
		}

		refreshTokenString, err := refreshToken.GenerateJWT()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
				Message: "Could not generate jwt string for refreshToken",
				Error:   err,
			})
		}
		accessTokenString, err := accessToken.GenerateJWT()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
				Message: "Could not generate jwt string for accessToken",
				Error:   err,
			})
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
			switch err.Error() {
			case errorCodes.InvalidPassword:
				return c.JSON(http.StatusUnauthorized, helpers.ResponseError{
					Message:    "Incorrect Password",
					Resolution: "Please use your correct password",
					Error:      err,
				})
			case errorCodes.InvalidEmail:
				return c.JSON(http.StatusUnauthorized, helpers.ResponseError{
					Message:    "Incorrect Email",
					Resolution: "Please use an email that corresponds to a user",
					Error:      err,
				})
			default:
				return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
					Message: "Database Error",
					Error:   err,
				})
			}
		}

		refreshToken, accessToken, err := con.userService.AddRefreshToken(user, c.RealIP())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
				Message: "Internal error while creating refreshToken",
				Error:   err,
			})
		}

		refreshTokenString, err := refreshToken.GenerateJWT()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
				Message: "Could not generate jwt string for refreshToken",
				Error:   err,
			})
		}
		accessTokenString, err := accessToken.GenerateJWT()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
				Message: "Could not generate jwt string for accessToken",
				Error:   err,
			})
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
			if err.Error() == errorCodes.TokenDoesNotExist {
				return c.JSON(http.StatusUnauthorized, helpers.ResponseError{
					Message:    "Your refresh token has been invalidated",
					Resolution: "Please login in to get a new refresh token",
				})
			}
			return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
				Message: err.Error(),
			})
		}

		accessTokenString, err := accessToken.GenerateJWT()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
				Message: "Could not generate tokenString",
				Error:   err,
			})
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
			if err.Error() == errorCodes.TokenDoesNotExist {
				return c.JSON(http.StatusUnauthorized, helpers.ResponseError{
					Message:    "Token doesn't exist",
					Resolution: "Use a refresh token that exists in database",
				})
			} else {
				return c.JSON(http.StatusInternalServerError, helpers.ResponseError{
					Message: "Error while deleting token",
					Error:   err,
				})
			}
		}
		return c.JSON(http.StatusOK, helpers.Json{
			"message": "successfully logged out user",
		})
	}
}
