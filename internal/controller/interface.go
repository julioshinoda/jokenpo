package controller

import (
	"context"

	"github.com/julioshinoda/jokenpo/internal/model"
)

type Matchcontroller interface {
	Evaluate(ctx context.Context, client model.MatchRequest) (string, error)
}
