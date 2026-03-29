package models

type SubmissionRequest struct {
	UserID     string `json:"user_id"`
	Timestamp  int64  `json:"timestamp"`
	Language   string `json:"language"`
	QuestionID string `json:"question_id"`
	Code       string `json:"code"`
}

type SubmissionResponse struct {
	Message      string `json:"message"`
	SubmissionID string `json:"submission_id,omitempty"`
}
