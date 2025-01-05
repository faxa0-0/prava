package storage

import (
	"github.com/faxa0-0/prava/models"
)

type QuizStorage interface {
	GenerateQuiz() (models.Quiz, error)
}
