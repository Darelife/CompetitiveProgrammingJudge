package models

type Submission struct {
	ID         int    `json:"id"`
	Code       string `json:"code"`
	Language   string `json:"language"`
	QuestionID string `json:"question_id"`
	Status     string `json:"status"`
	Output     string `json:"output"`
}
