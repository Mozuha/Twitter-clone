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

func ConnectDB() (*ent.Client, error) {
	log.Println("opening connection to dev database...")
	return newEntClient(false)
}

func ConnectTestDB() (*ent.Client, error) {
	log.Println("opening connection to test database...")
	return newEntClient(true)
}

func newEntClient(isTest bool) (*ent.Client, error) {
	host = "localhost"
	if isTest {
		port = os.Getenv("DB_TEST_PORT")
	} else {
		port = os.Getenv("DB_PORT")
	}
	user = os.Getenv("DB_USER")
	dbname = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	client, err := ent.Open("postgres", dbinfo)
	if err != nil {
		host = os.Getenv("DB_HOST")
		dbinfo = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
		client, err = ent.Open("postgres", dbinfo)
		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}
	}

	log.Println("connected to database")

	ctx := context.Background()

	// Auto migration; Enable UUID PK by passing WithGlobalUniqueID option, for GraphQL integration
	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if isTest {
		// Initially insert mock data for testing
		if err := InsertMockData(ctx, client); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("migration done")

	return client, err
}
