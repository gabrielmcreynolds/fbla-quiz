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

// Used to validate http requests
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// connect to db
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://gabriel:6VgmO49Db8KKk3Rv@testingmean.smyag.gcp.mongodb.net/<dbname>?retryWrites=true&w=majority"))
	if err != nil {
		log.Print(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Print(err)
	}

	defer client.Disconnect(ctx)
	log.Print("Connected to MongoDB")
	database := client.Database("fbla")
	// connected to db

	// start up http handler
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost:4200", "https://frontend-xqmdarcdtq-uc.a.run.app"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXCSRFToken},
	}))

	users.UserRoutes(e.Group("/users"), database)
	questions.QuestionRoutes(e.Group("/questions"), database)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
