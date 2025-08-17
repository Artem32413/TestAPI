package structs

// Warehouses представляет склад в системе
// swagger:model Warehouse
type Warehouses struct {
	// Уникальный идентификатор склада
	Id int `json:"id" example:"1"`

	// Внешний идентификатор склада
	WarehouseId string `json:"warehouseId" example:"WH-001"`

	// Физический адрес склада
	Addr string `json:"addr" example:"ул. Складская, д.1"`
}

type WarehousesSwagger struct {
	Addr string `json:"addr" example:"fghverv4446"`
}