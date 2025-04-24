package interfaces

import (
	"context"
	"graphql-poc/internal/models"
)

type TournamentsRepository interface {
	GetAll(ctx context.Context) ([]models.Tournament, error)
}
