package transport

import (
	"inventory/internal/handlers"
	"inventory/internal/config/database"
	"inventory/internal/database/postgres"
	"inventory/internal/service"
	"inventory/internal/transport/swagger"

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

	repo := postgres.NewInventoryRepository(db)
	svc := service.NewInventoryService(repo)
	handlers := handlers.NewInventoryHandlers(svc, log)

	mux := http.NewServeMux()

	swagger.AddSwaggerRoutes(mux)

	// Инвентаризация
	mux.HandleFunc("/inventory/price/", handlers.SetPrice)
	mux.HandleFunc("/inventory/updateQuantity/", handlers.UpdateInventory)
	mux.HandleFunc("/inventory/discount/", handlers.DiscountInventory)
	mux.HandleFunc("/inventory/goods/", handlers.ListOfGoods)
	mux.HandleFunc("/inventory/product/", handlers.ReceivingGoods)
	mux.HandleFunc("/inventory/count/", handlers.CountPrice)
	mux.HandleFunc("/inventory/purchase/", handlers.PurchaseProduct)
	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
