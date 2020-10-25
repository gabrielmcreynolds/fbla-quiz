package driver

import (
	"backend/helpers"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
	"time"
)

func getDatabaseConnection() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("dbString")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to MongoDB")
	database := client.Database("fbla")
	return database
}

func TestRepo_GetFiveQuestionsNoErr(t *testing.T) {
	db := getDatabaseConnection()
	repo := NewMongoRepository(db)
	questions, err := repo.GetFiveQuestions()
	if err != nil {
		t.Errorf("There was an error in repo.GetFiveQuestions: %v", err)
	}
	log.Printf("Questions: %v", questions)
	if len(questions) != 5 {
		t.Errorf("GetFiveQuestions: expected %d, actual %d", 5, len(questions))
	}
	helpers.Assert(t, len(questions) == 5)
	db.Client().Disconnect(context.Background())
}
