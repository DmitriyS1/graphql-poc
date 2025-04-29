package rest

import (
	"encoding/json"
	"github.com/jackc/pgx/v5/pgtype"
	"graphql-poc/internal/models"
	"graphql-poc/internal/repositories/interfaces"
	"net/http"
	"time"
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

func (th *TournamentHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var tournament models.Tournament
	if err := json.NewDecoder(r.Body).Decode(&tournament); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tournament.Created = pgtype.Date{Time: time.Now(), Valid: true}

	createdTournament, err := th.repository.Create(r.Context(), &tournament)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(createdTournament)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	return
}
