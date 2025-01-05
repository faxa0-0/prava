package services

import (
	"github.com/faxa0-0/prava/models"
	"github.com/faxa0-0/prava/storage"
)

type QuizService struct {
	storage storage.QuizStorage
}

func NewQuizService(storage storage.QuizStorage) *QuizService {
	return &QuizService{storage: storage}
}

func (srv *QuizService) GetQuiz() (models.Quiz, error) {
	return srv.storage.GenerateQuiz()
}
