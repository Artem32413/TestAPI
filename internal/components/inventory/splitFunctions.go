package inventory

import (
	"context"
	"fmt"
)

func (s *InventoryService) Exists(a int, value Inventory, value2 NewInventoryDiscount, value3 NewInventory, value4 WarehousePagination) bool {
	var exist bool

	if a == 1 {
		if err := s.Db.QueryRow(context.Background(), exists, value.Warehouse_id, value.Product_id).Scan(&exist); err != nil {
			s.Logger.Error("Ошибка в проверке на существование склада и товара")
		}
	} else if a == 2 {
		if err := s.Db.QueryRow(context.Background(), existsWarehouse, value.Warehouse_id).Scan(&exist); err != nil {
			s.Logger.Error("Ошибка в проверке на существование склада")
		}
	} else if a == 3 {
		if err := s.Db.QueryRow(context.Background(), existsWarehouse, value2.Warehouse_id).Scan(&exist); err != nil {
			s.Logger.Error("Ошибка в проверке на существование склада")
		}
	} else if a == 4 {
		if err := s.Db.QueryRow(context.Background(), existsWarehouse, value3.Warehouse_id).Scan(&exist); err != nil {
			s.Logger.Error("Ошибка в проверке на существование склада")
		}
	} else {
		if err := s.Db.QueryRow(context.Background(), existsWarehouse, value4.Warehouse_id).Scan(&exist); err != nil {
			s.Logger.Error("Ошибка в проверке на существование склада")
		}
	}

	return exist
}

func (s *InventoryService) GetAllAttributes(productIDs []string) (map[string]map[string]string, error) {
	if len(productIDs) == 0 {
		return make(map[string]map[string]string), nil
	}

	query := `SELECT product_id, key, value FROM product_key_values WHERE product_id = ANY($1)`

	rows, err := s.Db.Query(context.Background(), query, productIDs)
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
func convertAttributesToSlice(attrs map[string]string) []map[string]string {
	var result []map[string]string

	for key, value := range attrs {
		result = append(result, map[string]string{key: value})
	}
	return result
}

func Slices(count NewInventory) ([]string, []int) {
	productIDs := make([]string, len(count.Product))
	quantities := make([]int, len(count.Product))

	for i, p := range count.Product {
		productIDs[i] = p.Product_id
		quantities[i] = p.Quantity
	}

	return productIDs, quantities
}
