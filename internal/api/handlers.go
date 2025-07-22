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

func AddSwaggerRoutes(mux *http.ServeMux) {
	// Страница с Swagger UI (используем CDN)
	mux.HandleFunc("/docs/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<title>Swagger UI</title>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist/swagger-ui-bundle.js"></script>
<script>
  const ui = SwaggerUIBundle({
    url: "/docs/swagger.json",
    dom_id: '#swagger-ui',
  });
</script>
</body>
</html>`))
	})

	// Отдача файла swagger.json
	mux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})
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

	AddSwaggerRoutes(mux)
	// HealthCheck
	mux.HandleFunc("/api/health/", s.Health)

	// Склады
	mux.HandleFunc("/warehouses/add/", warehouse.AddingNewWarehouses)
	mux.HandleFunc("/warehouses/all/", warehouse.DisplayAllWarehouses)

	// Товары
	mux.HandleFunc("/products/add/", product.AddingNewProducts)
	mux.HandleFunc("/products/all/", product.DisplayAllProducts)
	mux.HandleFunc("/products/update/", product.UpdateProduct)

	// Инвентаризация
	mux.HandleFunc("/inventory/price/", inventory.SetPrice)
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
