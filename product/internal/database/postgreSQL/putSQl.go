package postgreSQL

import (
	"product/internal/model/structs"

	"context"
	"fmt"
)

func (d *DBService) UpdateProductSQL(ctx context.Context, products structs.Products) error {

	var exists bool
	err := d.db.QueryRow(ctx, productCheckUpd, products.ProductId).Scan(&exists)

	if err != nil {
		return fmt.Errorf("Ошибка в запросе: %v", err)
	}

	if !exists {
		return fmt.Errorf("Такого товара не существует: %v", err)
	}

	result, err := d.db.Exec(ctx, updateAProducts, products.Description, products.ProductId)

	if err != nil {
		return fmt.Errorf("Ошибка в запросе на обновление: %v", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Значение ниже 0: %d", rowsAffected)
	}

	for key, value := range products.KeyValue {
		if _, err := d.db.Exec(ctx, updateValue, value, products.ProductId, key); err != nil {
			return err
		}
	}

	return nil
}
