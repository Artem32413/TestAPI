package transport

import (
	"warehouses/internal/config/databaseConfig"
	"warehouses/internal/database/postgreSQL"
	"warehouses/internal/service"
	swaggerpkg "warehouses/internal/transport/swaggerPkg"
	"warehouses/internal/usecase"

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
	mux.HandleFunc("/warehouses/all/", handlers.DisplayAllWarehouses)
	mux.HandleFunc("/warehouses/add/", handlers.AddingNewWarehouses)
	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
