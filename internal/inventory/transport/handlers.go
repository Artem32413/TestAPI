package transport

import (
	swaggerpkg "apiGo/internal/analytics/transport/swaggerPkg"
	"apiGo/internal/inventory/appInventory"
	"apiGo/internal/inventory/config/databaseConfig"
	"apiGo/internal/inventory/database/postgreSQL"
	"apiGo/internal/inventory/service"

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
	handlers := appInventory.New(svc, log)
	mux := http.NewServeMux()

	swaggerpkg.AddSwaggerRoutes(mux)

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
