package errorCodes

import (
	"backend/helpers"
	"github.com/labstack/echo"
	"net/http"
)

type (
	Slug struct {
		ErrorCode  int
		Message    string
		Resolution string
	}
)

func (s Slug) Error() string {
	return s.Message
}

func (s Slug) Response(con *echo.Context) error {
	c := *con
	return c.JSON(s.ErrorCode, helpers.ResponseError{
		Message:    s.Message,
		Resolution: s.Resolution,
	})
}

func NewErrEmailAlreadyExists() *Slug {
	return &Slug{
		ErrorCode:  http.StatusConflict,
		Message:    SameEmailName,
		Resolution: "Try again with a different, unique email",
	}
}

func NewErrUserDoesNotExists() *Slug {
	return &Slug{
		ErrorCode:  http.StatusNotFound,
		Message:    UserDoesNotExist,
		Resolution: "User does not exist with the existing parameters",
	}
}

func NewErrInvalidPassword() *Slug {
	return &Slug{
		ErrorCode: http.StatusUnauthorized,
		Message:   InvalidPassword,
	}
}

func NewErrDatabaseIssue() *Slug {
	return &Slug{
		ErrorCode:  http.StatusInternalServerError,
		Message:    DatabaseError,
		Resolution: "Try again later",
	}
}

func NewErrInvalidEmail() *Slug {
	return &Slug{
		ErrorCode:  http.StatusNotFound,
		Message:    InvalidEmail,
		Resolution: "No user exists with that email",
	}
}

func NewErrInvalidRefreshToken() *Slug {
	return &Slug{
		ErrorCode:  http.StatusUnauthorized,
		Message:    InvalidRefreshToken,
		Resolution: "Try logging in again to get a new refresh token",
	}
}

func NewErrTokenDoesNotExist() *Slug {
	return &Slug{
		ErrorCode:  http.StatusBadRequest,
		Message:    TokenDoesNotExist,
		Resolution: "Token has been invalided. Try logging back in to get a new refresh token",
	}
}

func NewErrPasswordHashingFailure() *Slug {
	return &Slug{
		ErrorCode:  http.StatusInternalServerError,
		Message:    PasswordHashingFailure,
		Resolution: "There was an error in hashing the password",
	}
}

func NewErrTokenToJWT() *Slug {
	return &Slug{
		ErrorCode:  http.StatusInternalServerError,
		Message:    TokenToJWT,
		Resolution: "Try again later. Report this error to admin",
	}
}

const (
	SameEmailName          = "user with same emails already exists"
	UserDoesNotExist       = "user does not exist"
	InvalidPassword        = "invalid password"
	DatabaseError          = "database error"
	InvalidEmail           = "invalid email"
	InvalidRefreshToken    = "invalid refresh token"
	TokenDoesNotExist      = "token does not exist"
	PasswordHashingFailure = "passing hashing failure"
	TokenToJWT             = "could not parse token to jwt string"
)
