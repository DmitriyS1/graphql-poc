package models

import "github.com/jackc/pgx/v5/pgtype"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Hand string

const (
	RightHanded Hand = "right_handed"
	LeftHanded  Hand = "left_handed"
)

type Player struct {
	ID        int          `json:"id"`
	FirstName string       `json:"name"`
	LastName  string       `json:"last_name"`
	BirthDate pgtype.Date  `json:"birth_date"`
	Gender    Gender       `json:"sex"`
	Type      Hand         `json:"type"`
	Active    bool         `json:"active"`
	Created   pgtype.Date  `json:"created"`
	Updated   *pgtype.Date `json:"updated"`

	// relationships
	Stats       *PlayerStats  `json:"stats,omitempty"`
	Tournaments []*Tournament `json:"tournaments,omitempty"`
}
