package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"graphql-poc/internal/database"
	"graphql-poc/internal/models"
	"time"
)

type TournamentRepo struct {
	db database.DB
}

type GQLTournament struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Date          time.Time `json:"date"`
	PlayersAmount int       `json:"players_amount"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
}

func NewTournamentRepo(db database.DB) TournamentRepo {
	return TournamentRepo{db: db}
}

func (r TournamentRepo) GetAll(ctx context.Context) ([]models.Tournament, error) {
	rows, err := r.db.Pool.Query(ctx, "SELECT id, name, description, date, players_amount, created, updated FROM tournaments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Tournament])
}

func (r TournamentRepo) GetAllForGQL(ctx context.Context) ([]GQLTournament, error) {
	result, err := r.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var gqlTournaments []GQLTournament
	for _, tournament := range result {
		gqlTournament := GQLTournament{
			ID:            tournament.ID,
			Name:          tournament.Name,
			Description:   tournament.Description,
			Date:          tournament.Date.Time,
			PlayersAmount: tournament.PlayersAmount,
			Created:       tournament.Created.Time,
			Updated:       tournament.Updated.Time,
		}
		gqlTournaments = append(gqlTournaments, gqlTournament)
	}

	return gqlTournaments, nil
}

func (r TournamentRepo) Create(ctx context.Context, tournament *models.Tournament) (models.Tournament, error) {
	var id int
	err := r.db.Pool.QueryRow(ctx, "INSERT INTO tournaments (name, description, date, players_amount) VALUES ($1, $2, $3, $4) RETURNING id",
		tournament.Name, tournament.Description, tournament.Date, tournament.PlayersAmount).Scan(&id)
	if err != nil {
		return models.Tournament{}, err
	}

	tournament.ID = id
	return *tournament, nil
}
