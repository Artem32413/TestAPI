package postgreSQL

import (
	"apiGo/internal/inventory/database/postgreSQL/helpFunc"
	"apiGo/internal/inventory/model/structs"

	"context"

	"go.uber.org/zap"
)

func (d *DBService) SetPriceSQL(log *zap.Logger, ctx context.Context, price structs.Inventory) error {

	if exist := helpFunc.Exists(log, d.db, ctx, 1, price, structs.NewInventoryDiscount{}, structs.NewInventory{}, structs.WarehousePagination{}); exist {
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