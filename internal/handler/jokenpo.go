package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/julioshinoda/jokenpo/internal/controller"
	"github.com/julioshinoda/jokenpo/internal/model"
)

type Match struct {
	matchController controller.Matchcontroller
	validate        *validator.Validate
}

func NewMatch(mc controller.Matchcontroller, v *validator.Validate) Match {
	return Match{
		matchController: mc,
		validate:        v,
	}
}

func MatchHandlers(r chi.Router, m Match) {
	r.Post("/match", m.Evaluate)
}

func (m Match) Evaluate(w http.ResponseWriter, r *http.Request) {
	var req model.MatchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := m.validate.Struct(req); err != nil {
		http.Error(w, "Invalid request. Please,use valid values", http.StatusBadRequest)
		return
	}

	res, err := m.matchController.Evaluate(r.Context(), req)
	if err != nil {
		slog.Error(err.Error(), "method", "Evaluate")
		http.Error(w, "unexpected error", http.StatusInternalServerError)
		return

	}

	resp, err := json.Marshal(model.MatchResponse{
		Result: res,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(resp)); err != nil {
		// Handle the error, e.g., log it or return an error response
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
