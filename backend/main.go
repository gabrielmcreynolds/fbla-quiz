package main

import (
	"backend/questions"
	"backend/users"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// connect to db
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("dbString")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	log.Print("Connected to MongoDB")
	database := client.Database("fbla")

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())

	users.UserRoutes(e.Group("/users"), database)
	questions.QuestionRoutes(e.Group("/questions"), database)

	e.Logger.Fatal(e.Start(":8080"))
}
