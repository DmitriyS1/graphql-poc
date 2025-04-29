package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	gql "graphql-poc/handlers/graphql"
	"graphql-poc/internal/database"
	"graphql-poc/internal/repositories"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting graphql server on :8788")
	ctx := context.Background()

	dbAddr := os.Getenv("DB_ADDR")
	var db = database.NewPostgresDb(ctx, dbAddr)
	tournamentRepo := repositories.NewTournamentRepo(db)
	handler := gql.NewGraphQLHandler(tournamentRepo)
	schema, err := handler.InitSchema()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Only POST method is allowed")
			return
		}

		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName,omitempty"`
			Variables     map[string]interface{} `json:"variables,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error parsing request body: %v", err)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         *schema,
			OperationName:  params.OperationName,
			RequestString:  params.Query,
			VariableValues: params.Variables,
			Context:        r.Context(),
		})

		if len(result.Errors) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error executing query: %v", result.Errors)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error encoding response: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Query executed successfully")
	})

	http.ListenAndServe(":8788", nil)
}
