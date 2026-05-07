package queue

import "judge/internal/models"

var JobQueue = make(chan models.Submission, 100)
