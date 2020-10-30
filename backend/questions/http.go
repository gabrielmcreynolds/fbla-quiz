package questions

import (
	"backend/questions/controller"
	"backend/questions/driver"
	"backend/questions/usecases"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func QuestionRoutes(group *echo.Group, database *mongo.Database) {
	group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("accessSecret")),
	}))

	repo := driver.NewMongoRepository(database)
	service := usecases.NewService(repo)
	handler := controller.NewQuestionController(service)

	group.GET("", handler.GetFiveQuestions())
}
