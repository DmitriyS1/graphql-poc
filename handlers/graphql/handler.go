package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/jackc/pgx/v5/pgtype"
	"graphql-poc/internal/models"
	"graphql-poc/internal/repositories/interfaces"
)

type GraphQLHandler struct {
	repository interfaces.TournamentsRepository
}

func NewGraphQLHandler(repo interfaces.TournamentsRepository) *GraphQLHandler {
	return &GraphQLHandler{
		repository: repo,
	}
}

func (gh *GraphQLHandler) InitSchema() (*graphql.Schema, error) {
	var tournamentType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Tournament",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"players_amount": &graphql.Field{
				Type: graphql.Int,
			},
			"created": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})

	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"tournaments": &graphql.Field{
				Type: graphql.NewList(tournamentType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return gh.repository.GetAll(p.Context)
				},
			},
		},
	})

	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createTournament": &graphql.Field{
				Type: tournamentType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"date": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"players_amount": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name := p.Args["name"].(string)
					description := p.Args["description"].(string)
					date := p.Args["date"].(pgtype.Date)
					playersAmount := p.Args["players_amount"].(int)

					tournament := models.Tournament{
						Name:          name,
						Description:   description,
						Date:          date,
						PlayersAmount: playersAmount,
					}

					return gh.repository.Create(p.Context, &tournament)
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
