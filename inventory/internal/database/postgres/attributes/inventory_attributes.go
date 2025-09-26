package attributes

import (
	"context"
	"fmt"

	"inventory/internal/model/structs"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

var (
	existsWarehouse = `SELECT EXISTS(SELECT 1 FROM Inventory WHERE warehouseId = $1)`
	exists          = `SELECT EXISTS(SELECT 1 FROM Inventory WHERE warehouseId = $1 AND productId = $2)`
	query           = `SELECT productId, key, value FROM productKeyValues WHERE productId = ANY($1)`
)

func Exists(log *zap.Logger, db *pgx.Conn, ctx context.Context, a int, value structs.Inventory, value2 structs.NewInventoryDiscount, value3 structs.NewInventory, value4 structs.WarehousePagination) bool {
	var exist bool

	switch a {
	case 1:
		if err := db.QueryRow(ctx, exists, value.WarehouseId, value.ProductId).Scan(&exist); err != nil {
			log.Error("Ошибка в проверке на существование склада и товара")
		}
	case 2:
		if err := db.QueryRow(ctx, existsWarehouse, value.WarehouseId).Scan(&exist); err != nil {
			log.Error("Ошибка в проверке на существование склада")
		}
	case 3:
		if err := db.QueryRow(ctx, existsWarehouse, value2.WarehouseId).Scan(&exist); err != nil {
			log.Error("Ошибка в проверке на существование склада")
		}
	case 4:
		if err := db.QueryRow(ctx, existsWarehouse, value3.WarehouseId).Scan(&exist); err != nil {
			log.Error("Ошибка в проверке на существование склада")
		}
	default:
		if err := db.QueryRow(ctx, existsWarehouse, value4.WarehouseId).Scan(&exist); err != nil {
			log.Error("Ошибка в проверке на существование склада")
		}
	}

	return exist
}

func GetAllAttributes(db *pgx.Conn, productIDs []string) (map[string]map[string]string, error) {
	if len(productIDs) == 0 {
		return make(map[string]map[string]string), nil
	}

	rows, err := db.Query(context.Background(), query, productIDs)
	if err != nil {
		return nil, fmt.Errorf("Ошибка в поиске атрибутов: %w", err)
	}

	defer rows.Close()

	attrs := make(map[string]map[string]string)

	for rows.Next() {
		var productID, key, value string
		if err := rows.Scan(&productID, &key, &value); err != nil {
			return nil, fmt.Errorf("Ошибка в сканировании атрибутов: %w", err)
		}

		if _, exists := attrs[productID]; !exists {
			attrs[productID] = make(map[string]string)
		}

		attrs[productID][key] = value
	}

	return attrs, nil
}

func ConvertAttributesToSlice(attrs map[string]string) []map[string]string {
	var result []map[string]string

	for key, value := range attrs {
		result = append(result, map[string]string{key: value})
	}
	return result
}

func Slices(count structs.NewInventory) ([]string, []int) {
	productIDs := make([]string, len(count.Product))
	quantities := make([]int, len(count.Product))

	for i, p := range count.Product {
		productIDs[i] = p.ProductId
		quantities[i] = p.Quantity
	}

	return productIDs, quantities
}
