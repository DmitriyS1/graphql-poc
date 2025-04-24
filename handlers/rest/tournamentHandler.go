package rest

import (
	"encoding/json"
	"graphql-poc/internal/repositories/interfaces"
	"net/http"
)

type TournamentHandler struct {
	repository interfaces.TournamentsRepository
}

func NewTournamentHandler(repo interfaces.TournamentsRepository) *TournamentHandler {
	return &TournamentHandler{
		repository: repo,
	}
}

func (th *TournamentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	tournaments, err := th.repository.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(tournaments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
