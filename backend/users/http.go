package users

import (
	"backend/users/controller"
	"backend/users/driver"
	"backend/users/usecases"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func UserRoutes(group *echo.Group, database *mongo.Database) {
	// setup DI
	repo := driver.NewMongoRepository(database)
	service := usecases.NewService(repo)
	handler := controller.NewUserController(service)

	group.POST("/signup", handler.CreateUser())
	group.POST("/login", handler.Login())
	group.POST("/refresh", handler.Refresh())
	group.DELETE("/logout", handler.Logout(), middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("accessSecret")),
	}))
}
