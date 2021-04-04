package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tariqc80/oui-challenge/cmd/gqlgen/graph"
	"github.com/tariqc80/oui-challenge/cmd/gqlgen/graph/generated"
	"github.com/tariqc80/oui-challenge/internal/config"
	"github.com/tariqc80/oui-challenge/pkg/provider"
)

func main() {
	cfg := config.ParseEnv()

	// create new resolver and inject a new set provider with the config
	resolver := &graph.Resolver{
		Db:    provider.NewPg(cfg),
		Cache: provider.NewRedis(cfg),
	}

	defer resolver.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.GraphiqlPort)
	log.Fatal(http.ListenAndServe(":"+cfg.GraphiqlPort, nil))
}
