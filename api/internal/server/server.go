package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/humzamo/live-laugh-love/api/graph"
	"github.com/humzamo/live-laugh-love/api/graph/generated"
)

func StartServer() {
	fmt.Println("Starting server...")

	port := "8080"
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
