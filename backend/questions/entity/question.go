package entity

import "fmt"

// Simple Question struct that just handles the data, never processes.
type Question struct {
	Question      string        `json:"question" bson:"question"`
	Choices       []interface{} `json:"choices,omitempty" bson:"choices,omitempty"`
	CorrectChoice interface{}   `json:"correctChoice" bson:"correctChoice"`
}

func (q Question) String() string {
	return fmt.Sprintf("Question: %v\n\t Answer: %v\n", q.Question, q.CorrectChoice)
}
