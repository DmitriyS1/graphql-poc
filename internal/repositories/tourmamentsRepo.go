package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"graphql-poc/internal/database"
	"graphql-poc/internal/models"
)

type TournamentRepo struct {
	db database.DB
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
