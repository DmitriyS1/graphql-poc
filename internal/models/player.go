package models

import "github.com/jackc/pgx/v5/pgtype"

type Player struct {
	ID        int          `json:"id"`
	FirstName string       `json:"name"`
	LastName  string       `json:"last_name"`
	BirthDate pgtype.Date  `json:"birth_date"`
	Sex       bool         `json:"sex"`
	Type      string       `json:"type"` // "right_handed" or "left_handed"
	Active    bool         `json:"active"`
	Created   pgtype.Date  `json:"created"`
	Updated   *pgtype.Date `json:"updated"`
}
