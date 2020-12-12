package entity

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The full user that includes authentication & stats
type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	Name         string             `json:"name" bson:"name"`
	PasswordHash string             `json:"-" bson:"passwordHash"`
	TestsTaken   int                `json:"testsTaken" bson:"testsTaken"`
	TotalScores  int                `json:"totalScores" bson:"totalScores"`
	// in seconds
	TotalTime int `json:"totalTime" bson:"totalTime"`
}

// Authentication only portion of a user
type Authentication struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password" validate:"required"`
}

func (auth Authentication) String() string {
	return fmt.Sprintf("Email: %v, Name: %v, Password: %v", auth.Email, auth.Name, auth.Password)
}

func NewUserFromString(userString string) (*User, error) {
	user := new(User)
	err := json.Unmarshal([]byte(userString), user)
	return user, err
}
