package models

import "github.com/jackc/pgx/v5/pgtype"

type Tournament struct {
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Date          pgtype.Date  `json:"date"`
	PlayersAmount int          `json:"players_amount"`
	Created       pgtype.Date  `json:"created"`
	Updated       *pgtype.Date `json:"updated"`
}
