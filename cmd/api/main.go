package main

import (
	"database/sql"
	"log"

	"github.com/ctrlaltpat/skate-events/internal/env"
	"github.com/ctrlaltpat/skate-events/internal/handlers"
	"github.com/ctrlaltpat/skate-events/internal/repositories"
	"github.com/ctrlaltpat/skate-events/internal/services"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	port      int
	jwtSecret string
	db        *sql.DB
	handlers  *handlers.Handlers
}

func main() {
	db, err := sql.Open(("sqlite3"), "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repos := repositories.NewRepositories(db)
	services := services.NewServices(repos)
	handlers := handlers.NewHandlers(services)

	app := &application{
		port:      env.GetEnvInt("PORT", 5178),
		jwtSecret: env.GetEnvString("JWT_SECRET", "super-secret-key-88"),
		db:        db,
		handlers:  handlers,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
