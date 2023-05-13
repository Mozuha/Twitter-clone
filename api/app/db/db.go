package db

import (
	"api/ent"
	"api/ent/migrate"
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     string
	port     string
	user     string
	dbname   string
	password string
)

func NewEntClient(runningEnv string) (*ent.Client, error) {
	if runningEnv == "local" {
		host = "localhost"
	} else {
		host = os.Getenv("DB_HOST")
	}
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	dbname = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	client, err := ent.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	log.Println("connected to database")

	// Auto migration; Debug mode will print all SQL queries; Enable UUID PK by passing WithGlobalUniqueID option, for GraphQL integration
	if err := client.Debug().Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err
}
