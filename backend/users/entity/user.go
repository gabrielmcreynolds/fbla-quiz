package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	Name         string             `json:"name" bson:"name"`
	PasswordHash string             `json:"-" bson:"passwordHash"`
	TestsTaken   int                `json:"testsTaken" bson:"testsTaken"`
	TotalScores  int                `json:"totalScores" bson:"totalScores"`
	TotalTime    Duration           `json:"totalTime" bson:"totalTime"`
}

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

type Duration struct {
	time.Duration `json:"duration" bson:"duration"`
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}
