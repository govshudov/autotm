package handlers

import (
	"auth/internal/configs"
	"auth/internal/repository"
	"auth/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
)

const (
	baseURL = "/autotm"
	testURL = baseURL + "/test"
)

func Manager(logger *slog.Logger, cfg *configs.Config, repo *repository.PostgreSQLCartRepository) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//limitter
	//middleware

	r.Route(testURL, func(subRouter chi.Router) {
		testService := services.NewCartService(repo)
		testHandler := NewHTTPHandler(testService)
		testHandler.RegisterRoutes(subRouter)
	})
	return r
}
