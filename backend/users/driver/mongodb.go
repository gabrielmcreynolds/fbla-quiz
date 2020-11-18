package driver

import (
	"backend/errorCodes"
	"backend/users/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	// connectionString = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"
	connectionString = "mongodb://localhost:27017/?ssl=false"
)

type repo struct {
	Database *mongo.Database
}

func (r repo) AddUser(user *entity.User) (*entity.User, *errorCodes.Slug) {
	//TODO: add check to see if email is unique
	insertResult, err := r.Database.Collection("users").InsertOne(context.Background(), *user)
	if err != nil {
		log.Print(err)
		return nil, errorCodes.NewErrDatabaseIssue()
	}
	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r repo) FindUserByEmail(email string) (*entity.User, *errorCodes.Slug) {
	user := new(entity.User)
	err := r.Database.Collection("users").FindOne(context.Background(), bson.M{
		"email": email,
	}).Decode(user)

	if err != nil {
		log.Print(err)
		if err == mongo.ErrNoDocuments {
			return nil, errorCodes.NewErrInvalidEmail()
		}
		return nil, errorCodes.NewErrDatabaseIssue()
	}
	return user, nil
}

func (r repo) FindUserById(id *primitive.ObjectID) (*entity.User, *errorCodes.Slug) {
	user := new(entity.User)
	err := r.Database.Collection("users").FindOne(context.Background(), bson.M{
		"_id": id,
	}).Decode(user)

	if err != nil {
		log.Print(err)
		if err == mongo.ErrNoDocuments {
			return nil, errorCodes.NewErrUserDoesNotExists()
		}
		return nil, errorCodes.NewErrDatabaseIssue()
	}
	return user, nil
}

func (r repo) AddRefreshToken(token *entity.RefreshToken) (*entity.RefreshToken, *errorCodes.Slug) {
	result, err := r.Database.Collection("refreshTokens").InsertOne(context.Background(), token)
	if err != nil {
		log.Print(err)
		return nil, errorCodes.NewErrDatabaseIssue()
	}
	token.ID = result.InsertedID.(primitive.ObjectID)
	return token, nil
}

func (r repo) UpdateUser(user *entity.User) *errorCodes.Slug {
	filter := bson.D{{"_id", user.ID}}
	opts := options.FindOneAndReplace().SetReturnDocument(options.After)
	err := r.Database.Collection("users").FindOneAndReplace(context.Background(), filter, user, opts).Decode(&user)
	if err != nil {
		log.Print(err)
		if err == mongo.ErrNoDocuments {
			return errorCodes.NewErrUserDoesNotExists()
		}
	}
	return nil
}

func (r repo) FindRefreshTokenByID(id *primitive.ObjectID) (*entity.RefreshToken, *errorCodes.Slug) {
	refreshToken := new(entity.RefreshToken)
	err := r.Database.Collection("refreshTokens").FindOne(context.Background(), bson.D{{"_id", id}}).Decode(refreshToken)
	if err != nil {
		log.Print(err)
		if err == mongo.ErrNoDocuments {
			return nil, errorCodes.NewErrTokenDoesNotExist()
		}
		return nil, errorCodes.NewErrDatabaseIssue()
	}
	return refreshToken, nil
}

func (r repo) DeleteRefreshToken(id *primitive.ObjectID) *errorCodes.Slug {
	deletedToken := new(entity.RefreshToken)
	err := r.Database.Collection("refreshTokens").FindOneAndDelete(context.Background(), bson.D{{"_id", id}}).Decode(deletedToken)
	if err != nil {
		log.Print(err)
		if err == mongo.ErrNoDocuments {
			return errorCodes.NewErrTokenDoesNotExist()
		}
		return errorCodes.NewErrDatabaseIssue()
	}
	return nil
}

func NewMongoRepository(database *mongo.Database) Repository {
	return &repo{
		Database: database,
	}
}
