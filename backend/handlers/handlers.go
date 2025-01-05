package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/faxa0-0/prava/services"
)

type QuizHandler struct {
	service *services.QuizService
}

func NewQuizHandler(service *services.QuizService) *QuizHandler {
	return &QuizHandler{service: service}
}

func (handler *QuizHandler) GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	quiz, err := handler.service.GetQuiz()
	if err != nil {
		http.Error(w, "failed to generate quiz", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quiz)
}
