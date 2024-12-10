package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/julioshinoda/jokenpo/internal/controller"
	"github.com/julioshinoda/jokenpo/internal/handler"
	"gopkg.in/yaml.v2"
)

func main() {
	ctx := context.Background()

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte(`{"state":"up"}`)); err != nil {
			// Handle the error, e.g., log it or return an error response
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	rules := make(map[string][]string)

	yamlFile, err := os.ReadFile("./rules.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, rules)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	matchController := controller.NewMatch(rules)
	natchHandler := handler.NewMatch(matchController, validate)

	r.Group(func(r chi.Router) {
		//r.Use(authapp.Authenticator)
		handler.MatchHandlers(r, natchHandler)
	})

	slog.Info("Starting application on PORT: " + os.Getenv("SERVER_PORT"))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r); err != nil {
		log.Fatal(ctx, err.Error())
	}
}
