package api

import (
	"apiGo/internal/components"

	analytic "apiGo/internal/components/analytics"
	inventory "apiGo/internal/components/inventory"
	product "apiGo/internal/components/product"
	warehouse "apiGo/internal/components/warehouse"

	"net/http"
)

type InventoryService struct {
	*components.Settings
}

func AllHandles() *http.ServeMux {
	s, err := components.Set()

	if err != nil {
		s.Logger.Error(err.Error())
		return nil
	}

	analytic := &analytic.InventoryService{Settings: s}
	inventory := &inventory.InventoryService{Settings: s}
	product := &product.InventoryService{Settings: s}
	warehouse := &warehouse.InventoryService{Settings: s}

	mux := http.NewServeMux()

	// HealthCheck
	mux.HandleFunc("/api/health/", s.Health)

	// Склады
	mux.HandleFunc("/warehouses/add/", warehouse.AddingNewWarehouses)
	mux.HandleFunc("/warehouses/all/", warehouse.DisplayAllWarehouses)

	// Товары
	mux.HandleFunc("/products/add/", product.AddingNewProducts)
	mux.HandleFunc("/products/all/", product.DisplayAllProducts)
	mux.HandleFunc("/products/update/id", product.UpdateProduct)

	// Инвентаризация
	mux.HandleFunc("/inventory/connections/", inventory.Connection)
	mux.HandleFunc("/inventory/updateQuantity/", inventory.UpdateInventory)
	mux.HandleFunc("/inventory/discount/", inventory.DiscountInventory)
	mux.HandleFunc("/inventory/goods/", inventory.ListOfGoods)
	mux.HandleFunc("/inventory/product/", inventory.ReceivingGoods)
	mux.HandleFunc("/inventory/count/", inventory.CountPrice)
	mux.HandleFunc("/inventory/purchase/", inventory.PurchaseProduct)

	// Аналитика
	mux.HandleFunc("/analytics/", analytic.AnalyticsAll)
	mux.HandleFunc("/analytics/top/", analytic.Top)

	return mux
}
