package interfaces

import (
	"context"
	"graphql-poc/internal/models"
	"graphql-poc/internal/repositories"
)

type TournamentsRepository interface {
	GetAll(ctx context.Context) ([]models.Tournament, error)
	GetAllForGQL(ctx context.Context) ([]repositories.GQLTournament, error)
	Create(ctx context.Context, tournament *models.Tournament) (models.Tournament, error)
}
