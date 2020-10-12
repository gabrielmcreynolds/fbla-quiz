package usecases

import (
	"backend/errorCodes"
	"backend/users/entity"
)

type Service interface {
	CreateUser(user *entity.Authentication) (*entity.User, *errorCodes.Slug)
	Validate(auth *entity.Authentication) (*entity.User, *errorCodes.Slug)
	AddRefreshToken(user *entity.User, ip string) (*entity.RefreshToken, *entity.AccessToken, *errorCodes.Slug)
	RefreshToken(jwt string) (*entity.AccessToken, *errorCodes.Slug)
	RemoveRefreshToken(jwt string) *errorCodes.Slug
}
