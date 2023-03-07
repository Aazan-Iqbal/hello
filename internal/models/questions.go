// FIlename: internal/models/questions.go

package models

import (
	"context"
	"database/sql"
	"time"
)

// Lets model the questions table
type Question struct {
	QuestionID int64
	Body       string
	CreatedAt  time.Time
}

// Setup dependency injection of connction pool
type QuestionModel struct {
	DB *sql.DB
}

// Write sql code to access the database
//TODO

func (m *QuestionModel) Get() (*Question, error) {
	var q Question

	statement := `
				SELECT question_id, body
				FROM questions
				ORDER BY RANDOM()
				LIMIT 1
				`
	//context for when to end the query if it takes too long
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query the database using the connection pool m  
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.QuestionID, &q.Body)
	if err != nil {
		return nil, err
	}
	return &q, err
}
