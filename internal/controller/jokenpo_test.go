package controller

import (
	"context"
	"testing"

	"github.com/julioshinoda/jokenpo/internal/model"
)

func TestMatch_Evaluate(t *testing.T) {
	type fields struct {
		Rules map[string][]string
	}
	type args struct {
		ctx   context.Context
		match model.MatchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "player1 wins",
			fields: fields{
				Rules: map[string][]string{
					"rock":     []string{"scissors"},
					"scissors": []string{"paper"},
					"paper":    []string{"rock"},
				},
			},
			args: args{
				match: model.MatchRequest{
					Player1: "rock",
					Player2: "scissors",
				},
			},
			want: "player1",
		},
		{
			name: "player2 wins",
			fields: fields{
				Rules: map[string][]string{
					"rock":     []string{"scissors"},
					"scissors": []string{"paper"},
					"paper":    []string{"rock"},
				},
			},
			args: args{
				match: model.MatchRequest{
					Player1: "rock",
					Player2: "paper",
				},
			},
			want: "player2",
		},
		{
			name: "draw",
			fields: fields{
				Rules: map[string][]string{
					"rock":     []string{"scissors"},
					"scissors": []string{"paper"},
					"paper":    []string{"rock"},
				},
			},
			args: args{
				match: model.MatchRequest{
					Player1: "rock",
					Player2: "rock",
				},
			},
			want: "draw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Match{
				Rules: tt.fields.Rules,
			}
			got, err := u.Evaluate(tt.args.ctx, tt.args.match)
			if (err != nil) != tt.wantErr {
				t.Errorf("Match.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Match.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
