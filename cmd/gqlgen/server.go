package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tariqc80/oui-challenge/cmd/gqlgen/graph"
	"github.com/tariqc80/oui-challenge/cmd/gqlgen/graph/generated"
	"github.com/tariqc80/oui-challenge/internal/config"
	"github.com/tariqc80/oui-challenge/pkg/provider"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// create a config struct
	// TODO get values from env
	cfg := &config.Config{
		DatabaseName:     "postgres",
		DatabaseHost:     "postgres",
		DatabaseUser:     "postgres",
		DatabasePort:     "5432",
		DatabasePassword: "postgrespassword",
	}

	// create new resolver and inject a new set provider with the config
	resolver := &graph.Resolver{
		Provider: provider.NewSet(cfg),
	}

	defer resolver.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
