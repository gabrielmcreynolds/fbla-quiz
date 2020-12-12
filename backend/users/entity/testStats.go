package entity

type TestStats struct {
	Score int `json:"score" validate:"required"`
	// in seconds
	Time int `json:"time" validate:"required"`
}
