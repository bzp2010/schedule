package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bzp2010/schedule/internal/handler/graphql/generated"
	"github.com/bzp2010/schedule/internal/handler/graphql/resolvers"
)

// NewGraphQLHandler generates a handler for GraphQL requests
func NewGraphQLHandler() *handler.Server {
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolvers.Resolver{},
			},
		),
	)
	return server
}
