package postgreSQL

import (
	helpfunc "product/internal/database/postgreSQL/helpFunc"
	"product/internal/model/structs"

	"context"
	"fmt"
)

func (d *DBService) DisplayAllProductsSQL(ctx context.Context) ([]structs.Products, error) {
	productsRows, err := d.db.Query(ctx, displayProducts)
	if err != nil {
		return nil, fmt.Errorf("Запрос продуктов не удался: %w", err)
	}

	defer productsRows.Close()

	var products []structs.Products
	var productIDs []string

	for productsRows.Next() {
		var p structs.Products

		if err := productsRows.Scan(
			&p.ProductId,
			&p.Name,
			&p.Description,
			&p.Weight,
			&p.Barcode,
		); err != nil {
			return nil, fmt.Errorf("Сканирование продуктов не удалось: %w", err)
		}

		products = append(products, p)
		productIDs = append(productIDs, p.ProductId)
	}

	attrs, err := helpfunc.GetAllAttributes(d.db, ctx, productIDs)
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить атрибуты: %w", err)
	}

	for i := range products {
		if productAttrs, exists := attrs[products[i].ProductId]; exists {
			products[i].KeyValue = helpfunc.ConvertMapToSlice(productAttrs)
		}
	}

	return products, nil
}