// This acts as the file that constructs apis. (Connects handlers)
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/PulinduVR/ecom-go/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	productService := products.NewService()
	productHandler := products.NewHandler(productService)

	r.Get("/products", productHandler.ListProducts)

	return r

}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started on http://localhost%s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	//Can add logger and db driver here too.
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dbn string
}
