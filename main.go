package main

import (
	"database/sql"
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MudassirDev/barter/internal/web"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

const (
	DRIVER     string        = "libsql"
	EXPIRES_IN time.Duration = time.Hour * 1
)

var (
	DB_URL  string
	DB_CONN *sql.DB
	//go:embed db/migrations/*.sql
	embedMigrations embed.FS
	cfg             Config
)

func init() {
	godotenv.Load()
	log.Println("loading env variables!")

	port := os.Getenv("PORT")
	validateEnv(port, "PORT")

	dbURL := os.Getenv("DB_URL")
	validateEnv(dbURL, "DB_URL")
	DB_URL = dbURL

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	validateEnv(JWTSecretKey, "JWT_SECRET_KEY")

	environment := os.Getenv("ENV")

	log.Println("env variables loaded!")

	log.Println("setting up the server!")
	conn, err := sql.Open(DRIVER, dbURL)
	if err != nil {
		log.Fatalf("failed to make a connection with database: %v", err)
	}
	DB_CONN = conn

	handler := web.CreateMux(conn, JWTSecretKey, EXPIRES_IN, environment)
	cfg.handler = handler
	cfg.port = port

	log.Println("server setup done!")
}

// seperate for migrations
func init() {
	log.Println("making a connection to the database for migrations!")
	conn, err := sql.Open(DRIVER, DB_URL)
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

func main() {
	log.Println("starting the server!")
	defer DB_CONN.Close()
	srv := http.Server{
		Addr:    ":" + cfg.port,
		Handler: cfg.handler,
	}

	log.Printf("server is listening at port :%v!\n", cfg.port)
	log.Fatal(srv.ListenAndServe())
}

func validateEnv(variable, variableName string) {
	if variable == "" {
		log.Fatal("failed to load env variabe: ", variableName)
	}
}
