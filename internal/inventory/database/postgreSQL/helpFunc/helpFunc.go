package helpFunc

import (
	model "apiGo/internal/inventory/model/structs"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var (
	existsWarehouse = `SELECT EXISTS(SELECT 1 FROM Inventory WHERE warehouseId = $1)`
	exists          = `SELECT EXISTS(SELECT 1 FROM Inventory WHERE warehouseId = $1 AND productId = $2)`
)

func Exists(db *pgx.Conn, ctx context.Context, a int, value model.Inventory, value2 model.NewInventoryDiscount, value3 model.NewInventory, value4 model.WarehousePagination) bool {
	var exist bool

	if a == 1 {
		if err := db.QueryRow(ctx, exists, value.WarehouseId, value.ProductId).Scan(&exist); err != nil {
			s.Logger.Error("Ошибка в проверке на существование склада и товара")
		}
	} else if a == 2 {
		if err := db.QueryRow(ctx, existsWarehouse, value.WarehouseId).Scan(&exist); err != nil {
			s.Logger.Error("Ошибка в проверке на существование склада")
		}
	} else if a == 3 {
		if err := db.QueryRow(ctx, existsWarehouse, value2.WarehouseId).Scan(&exist); err != nil {
			s.Logger.Error("Ошибка в проверке на существование склада")
		}
	} else if a == 4 {
		if err := db.QueryRow(ctx, existsWarehouse, value3.WarehouseId).Scan(&exist); err != nil {
			s.Logger.Error("Ошибка в проверке на существование склада")
		}
	} else {
		if err := db.QueryRow(ctx, existsWarehouse, value4.WarehouseId).Scan(&exist); err != nil {
			s.Logger.Error("Ошибка в проверке на существование склада")
		}
	}

	return exist
}

func GetAllAttributes(db *pgx.Conn, productIDs []string) (map[string]map[string]string, error) {
	if len(productIDs) == 0 {
		return make(map[string]map[string]string), nil
	}

	query := `SELECT productId, key, value FROM product_key_values WHERE productId = ANY($1)`

	rows, err := db.Query(context.Background(), query, productIDs)
	if err != nil {
		return nil, fmt.Errorf("query attributes failed: %w", err)
	}

	defer rows.Close()

	attrs := make(map[string]map[string]string)

	for rows.Next() {
		var productID, key, value string
		if err := rows.Scan(&productID, &key, &value); err != nil {
			return nil, fmt.Errorf("scan attribute failed: %w", err)
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

func Slices(count model.NewInventory) ([]string, []int) {
	productIDs := make([]string, len(count.Product))
	quantities := make([]int, len(count.Product))

	for i, p := range count.Product {
		productIDs[i] = p.ProductId
		quantities[i] = p.Quantity
	}

	return productIDs, quantities
}