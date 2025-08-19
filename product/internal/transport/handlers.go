package transport

import (
	usecase "product/internal/usecase"
	"product/internal/config/databaseConfig"
	"product/internal/database/postgreSQL"
	"product/internal/service"
	swaggerpkg "product/internal/transport/swaggerPkg"
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

	// Товары
	mux.HandleFunc("/products/add/", handlers.AddingNewProducts)
	mux.HandleFunc("/products/all/", handlers.DisplayAllProducts)
	mux.HandleFunc("/products/update/", handlers.UpdateProduct)
	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
