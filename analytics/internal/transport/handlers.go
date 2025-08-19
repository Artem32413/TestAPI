package transport

import (
	"analytics/internal/usecase"
	"analytics/internal/config/databaseConfig"
	"analytics/internal/database/postgreSQL"
	"analytics/internal/service"
	swaggerpkg "analytics/internal/transport/swaggerPkg"

	"context"

	"net/http"

	"go.uber.org/zap"
)

func AllHandles(ctx context.Context, log *zap.Logger) *http.ServeMux {
	db, err := databaseConfig.ConstructorDB(ctx)

	if err != nil {
		log.Error(err.Error())
		return nil
	}

	repo := postgreSQL.New(db)
	svc := service.New(repo)
	handlers := usecase.New(svc, log)

	mux := http.NewServeMux()

	swaggerpkg.AddSwaggerRoutes(mux)

	// Аналитика
	mux.HandleFunc("/analytics/", handlers.AnalyticsAll)
	mux.HandleFunc("/analytics/top/", handlers.AnalyticsTop)
	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
