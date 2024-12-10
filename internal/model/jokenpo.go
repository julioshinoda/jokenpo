package model

type MatchRequest struct {
	Player1 string `json:"player1" validate:"oneof=scizor paper rock lizard spock" `
	Player2 string `json:"player2" validate:"oneof=scizor paper rock lizard spock"`
}

type MatchResponse struct {
	Result string `json:"result"`
}
