package postgreSQL

import (
	"inventory/internal/database/postgreSQL/helpFunc"
	"inventory/internal/model/structs"
	"fmt"

	"context"

	"go.uber.org/zap"
)

func (d *DBService) PurchaseProductSQL(log *zap.Logger, ctx context.Context, purchase structs.NewInventory) error {

	if exist := helpFunc.Exists(log, d.db, ctx, 4, structs.Inventory{}, structs.NewInventoryDiscount{}, purchase, structs.WarehousePagination{}); !exist {
		return fmt.Errorf("Склад не найден")
	}

	productIDs, quantities := helpFunc.Slices(purchase)

	var quantity int

	for _, productId := range productIDs {
		if err := d.db.QueryRow(ctx, quantityCheck, purchase.WarehouseId, productId).Scan(&quantity); err != nil {
			return fmt.Errorf("Ошибка проверки количества: %w", err)
		}
	}

	for i, el := range quantities {
		if el > quantity {
			return fmt.Errorf("Товар под номером %d отсутствует или это количество товара на складе отсутствует", i)
		}
	}
	for i, q := range quantities {
		pr := productIDs[i]

		if _, err := d.db.Exec(ctx, purchaseProduct, q, purchase.WarehouseId, pr); err != nil {
			return fmt.Errorf("Ошибка в списании товара со склада: %w", err)
		}

		if _, err := d.db.Exec(ctx, purchaseProductAnalytics, purchase.WarehouseId, pr, q); err != nil {
			return fmt.Errorf("Ошибка в записи данных в аналитику: %w", err)
		}
	}

	return nil
}