package graphql

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/graphql/generated"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/graphql/resolver"
)

type GraphQLServer struct {
	Port string
}

func NewGraphQLServer(port string) *GraphQLServer {
	return &GraphQLServer{
		Port: port,
	}
}

func (g *GraphQLServer) Start() {
	fmt.Println("Starting GraphQL server on port:", g.Port)

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})

	srv := handler.NewDefaultServer(schema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.ListenAndServe(":"+g.Port, nil)
}
