package postgres

import (
	"warehouses/internal/model/structs"

	"context"

	"github.com/google/uuid"
)

var (
	addingAWarehouse = `INSERT INTO Warehouses (warehouseId, addr) VALUES ($1, $2)`
)

func (d *DBService) AddingNewWarehousesSQL(ctx context.Context, warehouses structs.Warehouses) error {

	if d.db.IsClosed() {
		return d.db.Ping(ctx)
	}

	newUuid := uuid.New().String()

	warehouses.WarehouseId = newUuid

	if _, err := d.db.Exec(ctx, addingAWarehouse, warehouses.WarehouseId, warehouses.Addr); err != nil {
		return err
	}

	return nil
}
