package transport

import (
	"apiGo/internal/inventory/config/settings"
	"apiGo/internal/inventory/model/interfaces"

	"net/http"
)

type InventoryService struct {
	*settings.Settings
	interfaces.HandlersInventory
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
	s, err := settings.Set()

	if err != nil {
		s.Logger.Error(err.Error())
		return nil
	}

	var h interfaces.HandlersInventory

	h = &InventoryService{Settings: s}

	mux := http.NewServeMux()

	AddSwaggerRoutes(mux)

	// Инвентаризация
	mux.HandleFunc("/inventory/price/", h.SetPrice)
	mux.HandleFunc("/inventory/updateQuantity/", h.UpdateInventory)
	mux.HandleFunc("/inventory/discount/", h.DiscountInventory)
	mux.HandleFunc("/inventory/goods/", h.ListOfGoods)
	mux.HandleFunc("/inventory/product/", h.ReceivingGoods)
	mux.HandleFunc("/inventory/count/", h.CountPrice)
	mux.HandleFunc("/inventory/purchase/", h.PurchaseProduct)
	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
