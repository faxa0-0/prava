package postgres

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lib/pq"

	"github.com/faxa0-0/prava/models"
)

type PostgresQuizStorage struct {
	db *sql.DB
}

func NewPostgresQuizStorage(dsn string) (*PostgresQuizStorage, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresQuizStorage{db: db}, nil
}

func (storage *PostgresQuizStorage) GenerateQuiz() (models.Quiz, error) {
	query := `SELECT id, question, correct, incorrect, explanation, img_url FROM questions ORDER BY RANDOM() LIMIT 20;`
	rows, err := storage.db.Query(query)
	if err != nil {
		log.Print(err)
		return models.Quiz{}, err
	}
	defer rows.Close()

	var quiz models.Quiz
	for rows.Next() {
		var question models.Question
		err := rows.Scan(&question.ID, &question.Text, &question.Answer, pq.Array(&question.Options), &question.Explanation, &question.ImgURL)
		if err != nil {

			log.Print(err)
			return models.Quiz{}, err
		}
		quiz.Questions = append(quiz.Questions, question)
	}
	if err := rows.Err(); err != nil {

		log.Print(err)
		return models.Quiz{}, err
	}
	return quiz, nil
}
