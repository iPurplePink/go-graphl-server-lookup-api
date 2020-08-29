package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_"github.com/jinzhu/gorm/dialects/postgres"
	
	"github.com/go-graphl-server-lookup-api/config"
	"github.com/go-graphl-server-lookup-api/data/db"
	"github.com/go-graphl-server-lookup-api/graph"
	"github.com/go-graphl-server-lookup-api/graph/generated"
)

const defaultPort = "4030"

func init() {
	config.InitSettings()
}

func main() {

	db, err := db.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
