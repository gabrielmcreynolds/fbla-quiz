package driver

import (
	"backend/errorCodes"
	"backend/users/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	// Adds the entity.User to a repository
	// It returns the new entity.User with the ID field set correctly or an error if there was error.
	AddUser(user *entity.User) (*entity.User, *errorCodes.Slug)

	// Updates the user to the repository.
	// Assumes the ID field is set correctly in user
	UpdateUser(user *entity.User) *errorCodes.Slug

	// Returns a entity.User that contains the email given
	// If the user doesn't exist, then it will return nil, error("user doesn't exist")
	FindUserByEmail(email string) (*entity.User, *errorCodes.Slug)

	// Returns a entity.User that contains the id given
	// If the user doesn't exist, then it will return nil, error("user doesn't exist")
	FindUserById(id *primitive.ObjectID) (*entity.User, *errorCodes.Slug)

	// Adds the entity.RefreshToken to a repository
	// It returns the new entity.RefreshToken with the ID field set correctly or an error if there was error.
	AddRefreshToken(token *entity.RefreshToken) (*entity.RefreshToken, *errorCodes.Slug)

	// Returns a entity.RefreshToken that contains the email given
	// If the refreshToken doesn't exist, then it will return nil, error("refreshToken doesn't exist")
	FindRefreshTokenByID(id *primitive.ObjectID) (*entity.RefreshToken, *errorCodes.Slug)

	// Deletes the entity.RefreshToken
	// If the RefreshToken doesn't exist returns err("refreshToken doesn't exist")sxy
	DeleteRefreshToken(id *primitive.ObjectID) *errorCodes.Slug
}
