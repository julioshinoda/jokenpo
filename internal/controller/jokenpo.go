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

	/*if res := sort.SearchStrings(m.Rules[match.Player1], match.Player2); res < len(m.Rules[match.Player1]) {
		return "player1", nil
	}

	if res := sort.SearchStrings(m.Rules[match.Player2], match.Player1); res < len(m.Rules[match.Player2]) {
		return "player2", nil
	}*/

	// if m.Rules[match.Player1] == match.Player2 {
	// 	return "player1", nil
	// }

	// if m.Rules[match.Player2] == match.Player1 {
	// 	return "player2", nil
	// }x

	return "draw", nil
}
