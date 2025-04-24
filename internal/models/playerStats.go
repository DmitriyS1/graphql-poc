package models

import "github.com/jackc/pgx/v5/pgtype"

type PlayerStats struct {
	ID        int          `json:"id"`
	PlayerID  int          `json:"player_id"`
	GamesWon  int          `json:"games_won"`
	GamesLost int          `json:"games_lost"`
	Points    int          `json:"points"`
	Created   pgtype.Date  `json:"created"`
	Updated   *pgtype.Date `json:"updated"`
}
