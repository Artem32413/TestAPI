package postgreSQL

import (
	"apiGo/internal/product/config/databaseConfig"
	helpfunc "apiGo/internal/product/database/postgreSQL/helpFunc"
	"apiGo/internal/product/model/structs"

	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	displayProducts = `
	SELECT 
    p.productId,
    p.name,
    p.description,
    p.weight,
    p.barcode
FROM 
    products p
`
	displayKeyValue = `SELECT * FROM product_key_values WHERE productId = $1`
	addingAProducts = `INSERT INTO Products (productId, name, description, weight, barcode) VALUES ($1, $2, $3, $4, $5)`
	addingKeyValue  = `INSERT INTO product_key_values (productId, key, value) VALUES ($1, $2, $3)`
	deleteKeyValue  = `DELETE FROM product_key_values WHERE productId = $1`
	updateAProducts = `UPDATE Products SET description = $1 WHERE productId = $2`
	updateValue     = `UPDATE product_key_values SET value = $1 WHERE productId = $2 AND key = $3`
	productCheckUpd = `SELECT EXISTS(SELECT 1 FROM Products WHERE productId = $1)`
	productCheckIns = `SELECT EXISTS(SELECT 1 FROM Products WHERE name = $1)`
)

type DBService struct {
	db *pgx.Conn
}

func New(db *databaseConfig.PostgreSQL) *DBService {
	return &DBService{db: db.Db}
}

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
