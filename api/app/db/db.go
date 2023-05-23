package db

import (
	"app/ent"
	"app/ent/migrate"
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
	isTest   bool
)

func ConnectDB(runningEnv string) (*ent.Client, error) {
	if runningEnv == "docker" {
		host = os.Getenv("DB_HOST")
	} else {
		host = "localhost"
	}
	port = os.Getenv("DB_PORT")
	isTest = false

	log.Println("opening connection to dev database...")
	return newEntClient()
}

func ConnectTestDB(runningEnv string) (*ent.Client, error) {
	if runningEnv == "docker" {
		host = os.Getenv("DB_TEST_HOST")
		port = "5432"
	} else {
		host = "localhost"
		port = os.Getenv("DB_TEST_PORT")
	}
	isTest = true

	log.Println("opening connection to test database...")
	return newEntClient()
}

func newEntClient() (*ent.Client, error) {
	user = os.Getenv("DB_USER")
	dbname = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	client, err := ent.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	log.Println("connected to database")

	ctx := context.Background()

	// Auto migration
	// Enable Universal ID support by setting WithGlobalUniqueID option to true, for GraphQL integration (GraphQL requires that the object IDs are unique)
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
