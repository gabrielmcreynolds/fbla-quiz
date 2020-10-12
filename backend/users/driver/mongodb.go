package driver

import (
	"backend/errorCodes"
	"backend/users/entity"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	// connectionString = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"
	connectionString = "mongodb://localhost:27017/?ssl=false"
)

type repo struct {
	Database *mongo.Database
}

func (r repo) AddUser(user *entity.User) (*entity.User, error) {
	insertResult, err := r.Database.Collection("users").InsertOne(context.Background(), *user)
	if err != nil {
		log.Print(err)
		return nil, errors.New(errorCodes.DatabaseError)
	}
	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r repo) FindUserByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	err := r.Database.Collection("users").FindOne(context.Background(), bson.M{
		"email": email,
	}).Decode(user)

	if err != nil {
		log.Print(err)
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(errorCodes.InvalidEmail)
		}
		return nil, errors.New(errorCodes.DatabaseError)
	}
	return user, nil
}

func (r repo) FindUserById(id *primitive.ObjectID) (*entity.User, error) {
	user := new(entity.User)
	err := r.Database.Collection("users").FindOne(context.Background(), bson.M{
		"_id": id,
	})

	if err.Err() != nil {
		log.Print(err)
		log.Print(err.Err())
		if err.Err() == mongo.ErrNoDocuments {
			return nil, errors.New(errorCodes.UserDoesNotExist)
		}
		return nil, errors.New(errorCodes.DatabaseError)
	}
	return user, nil
}

func (r repo) AddRefreshToken(token *entity.RefreshToken) (*entity.RefreshToken, error) {
	result, err := r.Database.Collection("refreshTokens").InsertOne(context.Background(), token)
	if err != nil {
		log.Print(err)
		return nil, errors.New(errorCodes.DatabaseError)
	}
	token.ID = result.InsertedID.(primitive.ObjectID)
	return token, err
}

func (r repo) UpdateUser(user *entity.User) error {
	filter := bson.D{{"_id", user.ID.String()}}
	err := r.Database.Collection("users").FindOneAndReplace(context.Background(), filter, user).Decode(user)
	if err != nil {
		log.Print(err)
		if err == mongo.ErrNoDocuments {
			return errors.New(errorCodes.UserDoesNotExist)
		}
	}
	return nil
}

func (r repo) FindRefreshTokenByID(id *primitive.ObjectID) (*entity.RefreshToken, error) {
	refreshToken := new(entity.RefreshToken)
	err := r.Database.Collection("refreshTokens").FindOne(context.Background(), bson.D{{"_id", id}}).Decode(refreshToken)
	if err != nil {
		log.Print(err)
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(errorCodes.TokenDoesNotExist)
		}
		return nil, errors.New(errorCodes.DatabaseError)
	}

	return refreshToken, nil
}

func (r repo) DeleteRefreshToken(id *primitive.ObjectID) error {
	deletedToken := new(entity.RefreshToken)
	err := r.Database.Collection("refreshTokens").FindOneAndDelete(context.Background(), bson.D{{"_id", id}}).Decode(deletedToken)
	if err != nil {
		log.Print(err)
		if err == mongo.ErrNoDocuments {
			return errors.New(errorCodes.TokenDoesNotExist)
		}
		return errors.New(errorCodes.DatabaseError)
	}
	return nil
}

func NewMongoRepository(database *mongo.Database) Repository {
	return &repo{
		Database: database,
	}
}
