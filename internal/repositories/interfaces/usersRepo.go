package interfaces

type PlayersRepo interface {
	Get(id int, active *bool) (Player, error)
}
