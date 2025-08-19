package postgreSQL

import (
	"apiGo/internal/product/model/structs"

	"context"
	"fmt"

	"github.com/google/uuid"
)

func (d *DBService) AddingNewProductsSQL(ctx context.Context, products structs.Products) error {

	newUuid := uuid.New().String()

	products.ProductId = newUuid

	var exists bool
	err := d.db.QueryRow(ctx, productCheckIns, products.Name).Scan(&exists)

	if err != nil {
		return fmt.Errorf("Ошибка в сканировании: %w", err)
	}

	if exists {
		return fmt.Errorf("Товар уже существует")
	}

	if _, err := d.db.Exec(ctx, addingAProducts, products.ProductId, products.Name, products.Description, products.Weight, products.Barcode); err != nil {
		return fmt.Errorf("Ошибка с добавлением товара: %w", err)
	}

	if _, err := d.db.Exec(ctx, deleteKeyValue, products.ProductId); err != nil {
		return fmt.Errorf("Ошибка с удалением ключ значения: %w", err)
	}

	for _, m := range products.KeyValue {
		for key, value := range m {
			if _, err := d.db.Exec(ctx, addingKeyValue, products.ProductId, key, value); err != nil {
				return fmt.Errorf("Ошибка в добавлении ключ значения в бд: %w", err)
			}
		}
	}
	return nil
}
