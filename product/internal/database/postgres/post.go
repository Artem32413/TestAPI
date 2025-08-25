package postgres

import (
	"product/internal/model/structs"

	"context"
	"fmt"

	"github.com/google/uuid"
)

var(
	productCheckIns = `SELECT EXISTS(SELECT 1 FROM Products WHERE name = $1)`
	addingAProducts = `INSERT INTO Products (productId, name, description, weight, barcode) VALUES ($1, $2, $3, $4, $5)`
	deleteKeyValue  = `DELETE FROM ProductKeyValues WHERE productId = $1`
	addingKeyValue  = `INSERT INTO ProductKeyValues (productId, key, value) VALUES ($1, $2, $3)`
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
