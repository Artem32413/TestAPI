package postgreSQL

import (
	"apiGo/internal/inventory/database/postgreSQL/helpFunc"
	"apiGo/internal/inventory/model/structs"
	"fmt"

	"context"

	"go.uber.org/zap"
)

func (d *DBService) UpdateInventorySQL(log *zap.Logger, ctx context.Context, inventory structs.Inventory) error {

	if exist := helpFunc.Exists(log, d.db, ctx, 1, inventory, structs.NewInventoryDiscount{}, structs.NewInventory{}, structs.WarehousePagination{}); !exist {
		return fmt.Errorf("Ошибка в проверке на существование склада и товара")
	}

	if _, err := d.db.Exec(ctx, updateQuantity, inventory.Quantity, inventory.WarehouseId, inventory.ProductId); err != nil {
		return err
	}

	return nil
}

func (d *DBService) DiscountInventorySQL(log *zap.Logger, ctx context.Context, discount structs.NewInventoryDiscount) error {

	if exist := helpFunc.Exists(log, d.db, ctx, 3, structs.Inventory{}, discount, structs.NewInventory{}, structs.WarehousePagination{}); !exist {
		return fmt.Errorf("Склад не найден")
	}

	for _, el := range discount.ProductId {
		if _, err := d.db.Exec(ctx, creatingADiscount, discount.Discount, discount.WarehouseId, el); err != nil {
			return err
		}
	}

	return nil
}

func (d *DBService) CountPriceSQL(log *zap.Logger, ctx context.Context, count structs.NewInventory) (structs.SummingUp, error) {
	var result structs.SummingUp

	if exist := helpFunc.Exists(log, d.db, ctx, 4, structs.Inventory{}, structs.NewInventoryDiscount{}, count, structs.WarehousePagination{}); !exist {
		return result, fmt.Errorf("Склад с ID %s не найден", count.WarehouseId)
	}

	if len(count.Product) == 0 {
		return result, fmt.Errorf("Список товаров не может быть пустым")
	}

	productIDs, quantities := helpFunc.Slices(count)

	if err := d.db.QueryRow(
		ctx,
		listCount,
		productIDs,
		quantities,
		count.WarehouseId,
	).Scan(&result.Sum); err != nil {
		return result, fmt.Errorf("Ошибка расчета суммы: %w", err)
	}

	return result, nil
}