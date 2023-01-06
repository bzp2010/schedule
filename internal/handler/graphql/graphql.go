package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bzp2010/schedule/internal/handler/graphql/generated"
	"github.com/bzp2010/schedule/internal/handler/graphql/resolvers"
)

func NewGraphQLHandler() *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))
}
