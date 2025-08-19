package structs

type Analytics struct {
	WarehouseId string  `json:"warehouseId" example:"WH-001"` // Идентификатор склада
	ProductId   string  `json:"productId" example:"PRD-1001"` // Идентификатор товара
	SoldGoods   int     `json:"soldgoods" example:"50"`       // Количество проданных товаров
	TotalSum    float64 `json:"totalsum" example:"12500.50"`  // Общая сумма продаж
}

// TopAnalytics представляет данные топовых складов по выручке
// @Description Структура данных топовых складов по выручке
type TopAnalytics struct {
	Addr        string  `json:"addr" example:"ул. Ленина, 10"` // Адрес склада
	WarehouseId string  `json:"warehouseId" example:"WH-001"`  // Идентификатор склада
	TotalSum    float64 `json:"totalsum" example:"150000.75"`  // Общая выручка склада
}

type AnalyticsSwagger struct {
	WarehouseId string `json:"warehouseId" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"` // Идентификатор склада
}
