package structs

// WarehousePagination структура пагинации склада
// @Description Параметры пагинации для списка товаров на складе
type WarehousePagination struct {
	WarehouseId string `json:"warehouseId" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	Limit        int    `json:"limit" example:"10"` // Количество элементов на странице
	Offset       int    `json:"offset" example:"0"` // Смещение (номер страницы * limit)
}

// Inventory структура данных товара на складе
// @Description Основная информация о товаре на складе
type Inventory struct {
	WarehouseId string  `json:"warehouseId" example:"WH-001"`
	ProductId   string  `json:"productId" example:"PRD-1001"`
	Quantity     int     `json:"quantity,omitempty"`
	Price        float64 `json:"price,omitempty" example:"999.99"`
	Discount     float64 `json:"discount,omitempty"`
}

// NewInventory структура для создания/обновления инвентаря
// @Description Данные для создания или обновления инвентарной записи
type NewInventory struct {
	WarehouseId string              `json:"warehouseId" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	Product      []ProductInventory2 `json:"product"`
}

// NewInventoryDiscount структура для установки скидки
// @Description Данные для установки скидки на товары
type NewInventoryDiscount struct {
	WarehouseId string   `json:"warehouseId" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	ProductId   []string `json:"productId" example:"881670ed-8195-43d6-b678-3741bd3289ae,25f8f0ba-75f0-4caa-8648-27f6be3254b3"`
	Discount     float64  `json:"discount" example:"0.2"` // Размер скидки (0.2 = 20%)
}

// ProductInventory2 структура товара для операций
// @Description Упрощенная структура товара для операций
type ProductInventory2 struct {
	ProductId string `json:"productId" example:"881670ed-8195-43d6-b678-3741bd3289ae"`
	Quantity   int    `json:"quantity" example:"5"`
}

type InventorySwagger1 struct {
	WarehouseId string  `json:"warehouseId" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	ProductId   string  `json:"productId" example:"36c6ecf5-1e13-4cef-a591-bc8368fa0e60"`
	Price        float64 `json:"price,omitempty" example:"999.99"`
}

type InventorySwagger2 struct {
	WarehouseId string `json:"warehouseId" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	ProductId   string `json:"productId" example:"36c6ecf5-1e13-4cef-a591-bc8368fa0e60"`
	Quantity     int    `json:"quantity" example:"5"`
}

type InventorySwagger3 struct {
	WarehouseId string   `json:"warehouseId" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	ProductId   []string `json:"productId" example:"36c6ecf5-1e13-4cef-a591-bc8368fa0e60"`
	Quantity     int      `json:"quantity" example:"5"`
}

type InventorySwagger4 struct {
	WarehouseId string `json:"warehouseId" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	ProductId   string `json:"productId" example:"36c6ecf5-1e13-4cef-a591-bc8368fa0e60"`
}

type InventorySwagger5 struct {
	Sum float64 `json:"sum"`
}

type ListByWarehouse struct {
	ProductId string  `json:"productId"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Discount   float64 `json:"discount"`
}

type AllInformationAboutTheProduct struct {
	ProductId      string              `json:"productId"`
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	Characteristics []map[string]string `json:"characteristics,omitempty"`
	Barcode         string              `json:"barcode"`
	Price           float64             `json:"price"`
	Discount        float64             `json:"discount"`
	Quantity        int                 `json:"quantity"`
	Weight          float64             `json:"weight"`
}

type NewInventory2 struct {
	WarehouseID string             `json:"warehouseId"`
	Products    []ProductInventory `json:"product"`
}

type ProductInventory struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type SummingUp struct {
	Sum             float64           `json:"sum"`
	Characteristics map[string]string `json:"characteristics,omitempty"`
}
