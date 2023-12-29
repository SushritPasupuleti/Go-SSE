package routes

import (
	// "encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	// "context"
	// "server/env"
	// "server/handlers"
	"server/handlers"
	"server/helpers"
	// "server/models"
	// "strconv"
	// "github.com/rs/zerolog/log"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		indexHtml, err := os.ReadFile("./index.html")

		if err != nil {
			w.Write([]byte("Error reading index.html"))
			helpers.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		w.Write(indexHtml)
	})

	router.Group(func(r chi.Router) {
		router.Route("/stocks", func(r chi.Router) {
			r.Get("/{name}", handlers.GetStocks)
			r.Get("/stream/{name}", handlers.GetStocksStream)
		})
	})

	return router
}
