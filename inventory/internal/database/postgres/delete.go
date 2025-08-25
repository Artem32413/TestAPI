package postgres

import (
	"inventory/internal/database/postgres/attributes"
	"inventory/internal/model/structs"

	"context"
	"fmt"

	"go.uber.org/zap"
)

var (
	quantityCheck = `SELECT quantity FROM Inventory 
						WHERE warehouseId = $1 AND productId = $2`
	purchaseProduct = `UPDATE Inventory 
						SET quantity = quantity - $1 
						WHERE warehouseId = $2 AND productId = $3`
	purchaseProductAnalytics = `
								INSERT INTO analytics (warehouseId, productId, sold_goods, total_sum)
								SELECT 
									$1::text, 
									$2::text, 
									$3::integer,
									$3::integer * (
										SELECT price * (1 - COALESCE(discount, 0)) 
										FROM Inventory 
										WHERE productId = $2::text
									)
								ON CONFLICT (warehouseId, productId) 
								DO UPDATE SET
									sold_goods = analytics.sold_goods + EXCLUDED.sold_goods,
									total_sum = analytics.total_sum + EXCLUDED.total_sum`
)

func (d *DBService) PurchaseProductSQL(log *zap.Logger, ctx context.Context, purchase structs.NewInventory) error {

	if exist := attributes.Exists(log, d.db, ctx, 4, structs.Inventory{}, structs.NewInventoryDiscount{}, purchase, structs.WarehousePagination{}); !exist {
		return fmt.Errorf("Склад не найден")
	}

	productIDs, quantities := attributes.Slices(purchase)

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
