package transport

import (
	"context"
	
	"product/internal/config/database"
	"product/internal/database/postgres"
	"product/internal/handlers"
	"product/internal/service"
	"product/internal/transport/swagger"

	"net/http"

	"go.uber.org/zap"
)

func AllHandles(ctx context.Context, log *zap.Logger) *http.ServeMux {
	db, err := database.NewPostgreSQL(ctx)

	if err != nil {
		log.Error(err.Error())
		return nil
	}

	repo := postgres.NewProductRepository(db)
	svc := service.NewProductService(repo)
	handlers := handlers.NewProductHandler(svc, log)

	mux := http.NewServeMux()
	swagger.AddSwaggerRoutes(mux)

	// Товары
	mux.HandleFunc("/products/add/", handlers.AddingNewProducts)
	mux.HandleFunc("/products/all/", handlers.DisplayAllProducts)
	mux.HandleFunc("/products/update/", handlers.UpdateProduct)
	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
