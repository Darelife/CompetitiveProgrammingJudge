package worker

import (
	"fmt"

	"judge/internal/db"
	"judge/internal/judge"
	"judge/internal/queue"
)

func StartWorker() {
	go func() {
		fmt.Println("Worker started...")

		for job := range queue.JobQueue {
			status, output := judge.JudgeSubmission(
				job.Code,
				job.Language,
				job.QuestionID,
			)

			db.UpdateSubmission(
				job.ID,
				status,
				output,
			)

			fmt.Printf(
				"Processed submission %d: %s\n",
				job.ID,
				status,
			)
		}
	}()
}
