package models

type PlayerTournament struct {
	ID           int `json:"id"`
	PlayerID     int `json:"player_id"`
	TournamentID int `json:"tournament_id"`
	Place        int `json:"place"`
}
