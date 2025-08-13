package model

import "apiGo/internal/warehouses/config/settings"

// Warehouses представляет склад в системе
// swagger:model Warehouse
type Warehouses struct {
	// Уникальный идентификатор склада
	Id int `json:"id" example:"1"`

	// Внешний идентификатор склада
	Warehouse_id string `json:"warehouse_id" example:"WH-001"`

	// Физический адрес склада
	Addr string `json:"addr" example:"ул. Складская, д.1"`
}

type InventoryService struct {
	*settings.Settings
}