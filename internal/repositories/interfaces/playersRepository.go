package interfaces

import (
	"context"
	"graphql-poc/internal/models"
)

type PlayersRepo interface {
	Get(ctx context.Context, id int, active *bool) (models.Player, error)
	Create(ctx context.Context, player *models.Player) error
}
