package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/huaxk/hackernews/graph"
	"github.com/huaxk/hackernews/graph/generated"
	"github.com/huaxk/hackernews/repo"
)

func NewHandler(repo repo.Repository) http.Handler {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			Repository: repo,
		},
	}))
}

func NewPlaygroundHandler(endpoint string) http.Handler {
	return playground.Handler("GraphQL playground", endpoint)
}
