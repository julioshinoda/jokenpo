package controller

import (
	"context"

	"github.com/julioshinoda/jokenpo/internal/model"
)

type Match struct {
	Rules map[string][]string
}

func NewMatch(r map[string][]string) Match {
	return Match{
		Rules: r,
	}
}

func (m Match) Evaluate(ctx context.Context, match model.MatchRequest) (string, error) {

	for _, v := range m.Rules[match.Player1] {
		if v == match.Player2 {
			return "player1", nil
		}
	}

	for _, v := range m.Rules[match.Player2] {
		if v == match.Player1 {
			return "player2", nil
		}
	}

	return "draw", nil
}
