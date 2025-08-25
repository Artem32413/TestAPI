package transport

import (
	"warehouses/internal/config/database"
	"warehouses/internal/database/postgres"
	"warehouses/internal/handlers"
	"warehouses/internal/service"
	"warehouses/internal/transport/swagger"

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

	repo := postgres.NewWarehouseRepository(db)
	svc := service.NewWarehousesService(repo)
	handlers := handlers.NewWarehouseHandler(svc, log)

	mux := http.NewServeMux()

	swagger.AddSwaggerRoutes(mux)

	// Аналитика
	mux.HandleFunc("/warehouses/all/", handlers.DisplayAllWarehouses)
	mux.HandleFunc("/warehouses/add/", handlers.AddingNewWarehouses)
	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
