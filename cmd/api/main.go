package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/tracker-tv/actor-api/internal/handler"
	"github.com/tracker-tv/actor-api/internal/service"

	"github.com/jackc/pgx/v5"
)

type config struct {
	db struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

func main() {
	var cfg config

	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://test-user:test-password@localhost:5432/tracker-tv?sslmode=disable", "Postgres DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error("error opening database connection", err)
		os.Exit(1)
	}

	defer db.Close(context.Background())
	logger.Info("database connection pool established")

	s := service.New(db)

	mux := http.NewServeMux()
	h := handler.NewHandler(logger, s)
	mux.HandleFunc("/v1/actors", h.GetActors)
	mux.HandleFunc("/v1/actors/{id}", h.GetActor)

	logger.Info("starting server on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Error("error starting server", err)
		os.Exit(1)
	}

}

func openDB(cfg config) (*pgx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgx.Connect(context.Background(), cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
