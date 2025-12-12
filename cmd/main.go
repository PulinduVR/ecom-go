// Main entry point for the go backend. This composes all the dependencies and acts as the main entry point for the program execution.
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/PulinduVR/ecom-go/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	config := config{
		addr: ":8000",
		db: dbConfig{
			dbn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	//Database
	conn, err := pgx.Connect(ctx, config.db.dbn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	logger.Info("Database Connected", "dbn", config.db.dbn)

	api := application{
		config: config,
	}

	// h := api.mount()
	if err := api.run(api.mount()); err != nil {
		// log.Printf("Server has failed to start. Err : %s", err)
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
}
