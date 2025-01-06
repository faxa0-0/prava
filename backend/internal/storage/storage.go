package storage

import "github.com/faxa0-0/prava/internal/models"

type QuizStorage interface {
	GenerateQuiz() (models.Quiz, error)
}
