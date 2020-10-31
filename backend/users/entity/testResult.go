package entity

type TestResult struct {
	Score int      `json:"score" validate:"required"`
	Time  Duration `json:"time" validate:"required"`
}
