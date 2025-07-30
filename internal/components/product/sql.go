package product

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

var (
	displayProducts = `
	SELECT 
    p.product_id,
    p.name,
    p.description,
    p.weight,
    p.barcode
FROM 
    products p
`
	displayKeyValue = `SELECT * FROM product_key_values WHERE product_id = $1`
	addingAProducts = `INSERT INTO Products (product_id, name, description, weight, barcode) VALUES ($1, $2, $3, $4, $5)`
	addingKeyValue  = `INSERT INTO product_key_values (product_id, key, value) VALUES ($1, $2, $3)`
	deleteKeyValue  = `DELETE FROM product_key_values WHERE product_id = $1`
	updateAProducts = `UPDATE Products SET description = $1 WHERE product_id = $2`
	updateValue     = `UPDATE product_key_values SET value = $1 WHERE product_id = $2 AND key = $3`
	productCheckUpd = `SELECT EXISTS(SELECT 1 FROM Products WHERE product_id = $1)`
	productCheckIns = `SELECT EXISTS(SELECT 1 FROM Products WHERE name = $1)`
)

func (s *InventoryService) DisplayProducts() ([]Products, error) {
	productsRows, err := s.Db.Query(context.Background(), displayProducts)
	if err != nil {
		return nil, fmt.Errorf("Запрос продуктов не удался: %w", err)
	}

	defer productsRows.Close()

	var products []Products
	var productIDs []string

	for productsRows.Next() {
		var p Products

		if err := productsRows.Scan(
			&p.Product_id,
			&p.Name,
			&p.Description,
			&p.Weight,
			&p.Barcode,
		); err != nil {
			return nil, fmt.Errorf("Сканирование продуктов не удалось: %w", err)
		}

		products = append(products, p)
		productIDs = append(productIDs, p.Product_id)
	}

	attrs, err := s.GetAllAttributes(productIDs)
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить атрибуты: %w", err)
	}

	for i := range products {
		if productAttrs, exists := attrs[products[i].Product_id]; exists {
			products[i].KeyValue = convertMapToSlice(productAttrs)
		}
	}

	return products, nil
}
func convertMapToSlice(attrs map[string]string) []map[string]string {
	var result []map[string]string

	for key, value := range attrs {
		result = append(result, map[string]string{key: value})
	}

	return result
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
func (s *InventoryService) AdditionProducts(products Products) error {

	newUuid := uuid.New().String()

	products.Product_id = newUuid

	var exists bool
	err := s.Db.QueryRow(context.Background(), productCheckIns, products.Name).Scan(&exists)

	if err != nil {
		return fmt.Errorf("Ошибка в сканировании: %w", err)
	}

	if exists {
		return fmt.Errorf("Товар уже существует")
	}

	if _, err := s.Db.Exec(context.Background(), addingAProducts, products.Product_id, products.Name, products.Description, products.Weight, products.Barcode); err != nil {
		return fmt.Errorf("Ошибка с добавлением товара: %w", err)
	}

	if _, err := s.Db.Exec(context.Background(), deleteKeyValue, products.Product_id); err != nil {
		return fmt.Errorf("Ошибка с удалением ключ значения: %w", err)
	}

	for _, m := range products.KeyValue {
		for key, value := range m {
			if _, err := s.Db.Exec(context.Background(), addingKeyValue, products.Product_id, key, value); err != nil {
				return fmt.Errorf("Ошибка в добавлении ключ значения в бд: %w", err)
			}
		}

	}
	return nil
}

func (s *InventoryService) UpdateProd(products Products) error {

	var exists bool
	err := s.Db.QueryRow(context.Background(), productCheckUpd, products.Product_id).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return err
	}

	result, err := s.Db.Exec(context.Background(), updateAProducts, products.Description, products.Product_id)

	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return err
	}

	for key, value := range products.KeyValue {
		if _, err := s.Db.Exec(context.Background(), updateValue, value, products.Product_id, key); err != nil {
			return err
		}
	}

	return nil
}
