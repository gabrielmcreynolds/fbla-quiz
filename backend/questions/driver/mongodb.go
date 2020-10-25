package driver

import (
	"backend/errorCodes"
	"backend/questions/entity"
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	Database *mongo.Database
}

func (r repo) GetFiveQuestions() ([]*entity.Question, *errorCodes.Slug) {
	groupStage := []bson.D{bson.D{{"$sample", bson.D{{"size", 5}}}}}
	cursor, err := r.Database.Collection("questions").Aggregate(context.Background(), groupStage)
	if err != nil {
		log.Error(err)
		return nil, errorCodes.NewErrDatabaseIssue()
	}
	var results []*entity.Question
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Error(err)
		return nil, errorCodes.NewErrDatabaseIssue()
	}
	return results, nil
}

func NewMongoRepository(database *mongo.Database) Repository {
	return &repo{
		Database: database,
	}
}
