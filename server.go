package main

import (
	"apiGrapqlEntgo/ent"
	"apiGrapqlEntgo/graph"
	"apiGrapqlEntgo/graph/generated"
	"context"
	"database/sql"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	pgDriver = "pgx" //This is the name of the postgres driver registered by jackc/pgx
	port     = "8080"
)

// getClient creates a new *ent.Client driver
func getClient(connStr string) (*ent.Client, error) {
	// Open Database
	db, err := sql.Open(pgDriver, connStr)
	if err != nil {
		return nil, err
	}

	// Create driver and return
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver)), nil
}

func main() {
	// Create client
	client, err := getClient("host=localhost port=5432 user=testuser dbname=test_fullapi password=testpswd")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating ORM resources: %v", err)
	}

	// Create resolver
	resolver := &graph.Resolver{
		Cli: client,
	}

	// Init service
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
