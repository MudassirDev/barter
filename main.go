package main

import (
	"database/sql"
	"embed"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var (
	DB_URL string
	//go:embed db/migrations/*.sql
	embedMigrations embed.FS
)

func init() {
	godotenv.Load()
	log.Println("loading env variables!")

	port := os.Getenv("PORT")
	validateEnv(port, "PORT")

	dbURL := os.Getenv("DB_URL")
	validateEnv(dbURL, "DB_URL")
	DB_URL = dbURL

	log.Println("env variables loaded!")
}

// seperate for migrations
func init() {
	log.Println("making a connection to the database for migrations!")
	conn, err := sql.Open("libsql", DB_URL)
	if err != nil {
		log.Fatalf("failed to make a connection with database: %v", err)
	}
	defer conn.Close()
	log.Println("database connection formed successfully!")

	log.Println("running migrations!")

	goose.SetDialect("turso")
	goose.SetBaseFS(embedMigrations)

	if err := goose.Up(conn, "db/migrations"); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("migrations ran successfully!")
}

func main() {}

func validateEnv(variable, variableName string) {
	if variable == "" {
		log.Fatal("failed to load env variabe: ", variableName)
	}
}
