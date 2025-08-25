package transport

import (
	"analytics/internal/config/database"
	"analytics/internal/database/postgres"
	"analytics/internal/handlers"
	"analytics/internal/service"
	"analytics/internal/transport/swagger"

	"context"
	"net/http"

	"go.uber.org/zap"
)

func AllHandles(ctx context.Context, log *zap.Logger) *http.ServeMux {
	db, err := database.NewPostgreSQL(ctx)

	if err != nil {
		log.Error(err.Error())
		return nil
	}

	repo := postgres.NewAnalyticsRepository(db)
	svc := service.NewAnalyticsService(repo)
	handlers := handlers.NewAnalyticsHandler(svc, log)

	mux := http.NewServeMux()

	swagger.AddSwaggerRoutes(mux)

	mux.HandleFunc("/analytics/", handlers.AnalyticsAll)
	mux.HandleFunc("/analytics/top/", handlers.AnalyticsTop)
	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
