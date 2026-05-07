package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() error {
	var err error

	DB, err = sql.Open("sqlite3", "judge.db")
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS submissions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		language TEXT,
		question_id TEXT,
		status TEXT,
		output TEXT
	)
	`

	_, err = DB.Exec(query)
	return err
}

func CreateSubmission(code, language, questionID string) (int64, error) {
	result, err := DB.Exec(
		`INSERT INTO submissions
		(code, language, question_id, status, output)
		VALUES (?, ?, ?, ?, ?)`,
		code,
		language,
		questionID,
		"PENDING",
		"",
	)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func UpdateSubmission(id int, status, output string) error {
	_, err := DB.Exec(
		"UPDATE submissions SET status=?, output=? WHERE id=?",
		status,
		output,
		id,
	)

	return err
}

func GetSubmission(id int) (map[string]interface{}, error) {
	row := DB.QueryRow(
		"SELECT id, language, question_id, status, output FROM submissions WHERE id=?",
		id,
	)

	var submissionID int
	var language string
	var questionID string
	var status string
	var output string

	err := row.Scan(
		&submissionID,
		&language,
		&questionID,
		&status,
		&output,
	)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":          submissionID,
		"language":    language,
		"question_id": questionID,
		"status":      status,
		"output":      output,
	}, nil
}
