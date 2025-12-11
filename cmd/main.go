// Main entry point for the go backend. This composes all the dependencies and acts as the main entry point for the program execution.
package main

import (
	"log/slog"
	"os"
)

func main() {
	config := config{
		addr: ":8000",
		db:   dbConfig{},
	}

	api := application{
		config: config,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	// h := api.mount()
	if err := api.run(api.mount()); err != nil {
		// log.Printf("Server has failed to start. Err : %s", err)
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
}
