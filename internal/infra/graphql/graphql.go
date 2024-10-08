package graphql

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/graphql/generated"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/graphql/resolver"
	"github.com/renanmav/GoExpert-CleanArch/internal/usecase"
)

type GraphQLServer struct {
	Port                 string
	CreateOrderUseCase   usecase.CreateOrderUseCase
	ReadAllOrdersUseCase usecase.ReadAllOrdersUseCase
	ReadOrderByIdUseCase usecase.ReadOrderByIdUseCase
}

func NewGraphQLServer(
	port string,
	createOrderUseCase usecase.CreateOrderUseCase,
	readAllOrdersUseCase usecase.ReadAllOrdersUseCase,
	readOrderByIdUseCase usecase.ReadOrderByIdUseCase,
) *GraphQLServer {
	return &GraphQLServer{
		Port:                 port,
		CreateOrderUseCase:   createOrderUseCase,
		ReadAllOrdersUseCase: readAllOrdersUseCase,
		ReadOrderByIdUseCase: readOrderByIdUseCase,
	}
}

func (g *GraphQLServer) Start() {
	fmt.Println("Starting GraphQL server on port:", g.Port)

	r := resolver.Resolver{
		CreateOrderUseCase:   g.CreateOrderUseCase,
		ReadAllOrdersUseCase: g.ReadAllOrdersUseCase,
		ReadOrderByIdUseCase: g.ReadOrderByIdUseCase,
	}

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &r})

	srv := handler.NewDefaultServer(schema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.ListenAndServe(":"+g.Port, nil)
}
