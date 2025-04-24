package main

import (
	"context"
	"fmt"
	"graphql-poc/handlers/rest"
	"graphql-poc/internal/database"
	"graphql-poc/internal/repositories"
	"net/http"
	"os"
)

func main() {
	router := http.NewServeMux()
	fmt.Println("Starting server on :8787")
	ctx := context.Background()
	dbAddr := os.Getenv("DB_ADDR")
	var db = database.NewPostgresDb(ctx, dbAddr)
	tournamentRepo := repositories.NewTournamentRepo(db)

	tournamentHandler := rest.NewTournamentHandler(tournamentRepo)

	router.HandleFunc("GET /api/v1/tournaments", tournamentHandler.GetAll)
	router.HandleFunc("POST /api/v1/tournaments", tournamentHandler.Create)

	err := http.ListenAndServe(":8787", router)
	if err != nil {
		panic(err)
	}
}
