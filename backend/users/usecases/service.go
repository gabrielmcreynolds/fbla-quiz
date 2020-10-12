package usecases

import (
	"backend/users/entity"
)

type Service interface {
	CreateUser(user *entity.Authentication) (*entity.User, error)
	Validate(auth *entity.Authentication) (*entity.User, error)
	AddRefreshToken(user *entity.User, ip string) (*entity.RefreshToken, *entity.AccessToken, error)
	RefreshToken(jwt string) (*entity.AccessToken, error)
	RemoveRefreshToken(jwt string) error
}
