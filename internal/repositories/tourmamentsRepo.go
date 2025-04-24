package repositories

import (
	"context"
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

	var tournaments []models.Tournament
	for rows.Next() {
		var tournament models.Tournament
		if err := rows.Scan(&tournament); err != nil {
			return nil, err
		}
		tournaments = append(tournaments, tournament)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tournaments, nil
}
