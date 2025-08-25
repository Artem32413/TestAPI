package postgres

import (
	"inventory/internal/database/postgres/attributes"
	"inventory/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

var (
	priceUpdate = `UPDATE Inventory SET price = $1 WHERE warehouseId = $2 AND productId = $3`
	priceInsert = `INSERT INTO Inventory (warehouseId, productId, quantity, price, discount) VALUES ($1, $2, $3, $4, $5)`
)

func (d *DBService) SetPriceSQL(log *zap.Logger, ctx context.Context, price structs.Inventory) error {

	if exist := attributes.Exists(log, d.db, ctx, 1, price, structs.NewInventoryDiscount{}, structs.NewInventory{}, structs.WarehousePagination{}); exist {
		if _, err := d.db.Exec(ctx, priceUpdate, price.Price, price.WarehouseId, price.ProductId); err != nil {
			return err
		}
	} else {
		if _, err := d.db.Exec(ctx, priceInsert, price.WarehouseId, price.ProductId, 0, price.Price, 0); err != nil {
			return err
		}
	}

	return nil
}
