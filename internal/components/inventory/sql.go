package inventory

import (
	"context"
	"fmt"
)

var (
	exists          = `SELECT EXISTS(SELECT 1 FROM Inventory WHERE warehouse_id = $1 AND product_id = $2)`
	existsWarehouse = `SELECT EXISTS(SELECT 1 FROM Inventory WHERE warehouse_id = $1)`
	priceInsert     = `INSERT INTO Inventory (warehouse_id, product_id, quantity, price, discount) VALUES ($1, $2, $3, $4, $5)`
	priceUpdate     = `UPDATE Inventory SET price = $1 WHERE warehouse_id = $2 AND product_id = $3`
	updateQuantity  = `UPDATE Inventory 
						SET quantity = quantity + $1
						WHERE warehouse_id = $2 AND product_id = $3`
	creatingADiscount = `UPDATE Inventory SET discount = $1 WHERE warehouse_id = $2 AND product_id = $3`
	listOfGoods       = `SELECT 
						p.product_id,
						p.name,
						i.price,
						ROUND(i.price * (1 - i.discount/100), 2) as discounted_price
						FROM inventory i
						JOIN products p ON i.product_id = p.product_id
						WHERE i.warehouse_id = $1
						LIMIT $2 OFFSET $3`
	listInventory = `SELECT 
						COALESCE(quantity, 0) AS quantity,
						COALESCE(price, 0) AS price,
						COALESCE(discount, 0) AS discount 
						FROM Inventory 
						WHERE warehouse_id = $1 AND product_id = $2`
	oneProduct = `SELECT 
					p.product_id,
					p.name,
					p.description,
					p.weight,
					p.barcode
						FROM 
					Products p
					WHERE product_id = $1`
	listCount = `SELECT COALESCE(SUM(i.price * p.quantity * (1 - i.discount/100)), 0)
					FROM unnest($1::text[], $2::int[]) AS p(product_id, quantity)
					JOIN inventory i ON i.product_id = p.product_id AND i.warehouse_id = $3
					WHERE i.quantity >= p.quantity`
	quantityCheck = `SELECT quantity FROM Inventory 
						WHERE warehouse_id = $1 AND product_id = $2`
	purchaseProduct = `UPDATE Inventory 
						SET quantity = quantity - $1 
						WHERE warehouse_id = $2 AND product_id = $3`
)

type ListByWarehouse struct {
	Product_id string  `json:"product_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Discount   float64 `json:"discount"`
}

type AllInformationAboutTheProduct struct {
	Product_id      string              `json:"product_id"`
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	Characteristics []map[string]string `json:"characteristics,omitempty"`
	Barcode         string              `json:"barcode"`
	Price           float64             `json:"price"`
	Discount        float64             `json:"discount"`
	Quantity        int                 `json:"quantity"`
	Weight          float64             `json:"weight"`
}

type NewInventory2 struct {
	WarehouseID string             `json:"warehouse_id"`
	Products    []ProductInventory `json:"product"`
}

type ProductInventory struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type SummingUp struct {
	Sum             float64           `json:"sum"`
	Characteristics map[string]string `json:"characteristics,omitempty"`
}

func (s *InventoryService) SetPriceDB(price Inventory) error {

	exist := s.Exists(1, price, NewInventoryDiscount{}, NewInventory{}, WarehousePagination{})

	if exist {
		if _, err := s.Db.Exec(context.Background(), priceUpdate, price.Price, price.Warehouse_id, price.Product_id); err != nil {
			return err
		}
	} else {
		if _, err := s.Db.Exec(context.Background(), priceInsert, price.Warehouse_id, price.Product_id, 0, price.Price, 0); err != nil {
			return err
		}
	}

	return nil
}

func (s *InventoryService) UpdateQuantity(inventory Inventory) error {

	exist := s.Exists(1, inventory, NewInventoryDiscount{}, NewInventory{}, WarehousePagination{})

	if !exist {
		return fmt.Errorf("Товар не найден")
	}

	if _, err := s.Db.Exec(context.Background(), updateQuantity, inventory.Quantity, inventory.Warehouse_id, inventory.Product_id); err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) CreatingADiscount(discount NewInventoryDiscount) error {

	exist := s.Exists(3, Inventory{}, discount, NewInventory{}, WarehousePagination{})

	if !exist {
		return fmt.Errorf("Склад не найден")
	}

	for _, el := range discount.Product_id {
		if _, err := s.Db.Exec(context.Background(), creatingADiscount, discount.Discount, discount.Warehouse_id, el); err != nil {
			return err
		}
	}

	return nil
}

func (s *InventoryService) ListProductsByWarehouse(warehouse WarehousePagination) ([]ListByWarehouse, error) {

	exist := s.Exists(5, Inventory{}, NewInventoryDiscount{}, NewInventory{}, warehouse)

	if !exist {
		return nil, fmt.Errorf("склад с ID %s не найден", warehouse.Warehouse_id)
	}

	rows, err := s.Db.Query(
		context.Background(),
		listOfGoods,
		warehouse.Warehouse_id,
		warehouse.Limit,
		warehouse.Offset,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}

	defer rows.Close()

	var products []ListByWarehouse

	for rows.Next() {
		var p ListByWarehouse

		err := rows.Scan(
			&p.Product_id,
			&p.Name,
			&p.Price,
			&p.Discount,
		)

		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования данных: %w", err)
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при обработке результатов: %w", err)
	}

	return products, nil
}

func (s *InventoryService) ListProduct(product Inventory) (AllInformationAboutTheProduct, error) {

	exist := s.Exists(1, product, NewInventoryDiscount{}, NewInventory{}, WarehousePagination{})

	if !exist {
		return AllInformationAboutTheProduct{}, fmt.Errorf("Товар не найден")
	}

	var n AllInformationAboutTheProduct

	err := s.Db.QueryRow(context.Background(), listInventory, product.Warehouse_id, product.Product_id).Scan(&n.Price, &n.Discount, &n.Quantity)
	if err != nil {
		return AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка с запросом остатка %w", err)
	}

	err = s.Db.QueryRow(context.Background(), oneProduct, product.Product_id).Scan(&n.Product_id, &n.Name, &n.Description, &n.Weight, &n.Barcode)
	if err != nil {
		return AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка с запросом товара %w", err)
	}

	attrs, err := s.GetAllAttributes([]string{n.Product_id})
	if err != nil {
		return n, fmt.Errorf("ошибка получения характеристик: %w", err)
	}

	if productAttrs, exists := attrs[n.Product_id]; exists {
		n.Characteristics = convertAttributesToSlice(productAttrs)
	}

	return n, nil
}

func (s *InventoryService) ListCount(count NewInventory) (SummingUp, error) {
	var result SummingUp

	exist := s.Exists(4, Inventory{}, NewInventoryDiscount{}, count, WarehousePagination{})

	if !exist {
		return result, fmt.Errorf("склад с ID %s не найден", count.Warehouse_id)
	}

	if len(count.Product) == 0 {
		return result, fmt.Errorf("список товаров не может быть пустым")
	}

	productIDs, quantities := Slices(count)

	err := s.Db.QueryRow(
		context.Background(),
		listCount,
		productIDs,
		quantities,
		count.Warehouse_id,
	).Scan(&result.Sum)

	if err != nil {
		return result, fmt.Errorf("ошибка расчета суммы: %w", err)
	}

	return result, nil
}

func (s *InventoryService) Purchase(purchase NewInventory) error {

	exist := s.Exists(4, Inventory{}, NewInventoryDiscount{}, purchase, WarehousePagination{})

	if !exist {
		return fmt.Errorf("Склад не найден")
	}

	productIDs, quantities := Slices(purchase)

	var quantity int

	for _, product_id := range productIDs {
		err := s.Db.QueryRow(context.Background(), quantityCheck, purchase.Warehouse_id, product_id).Scan(&quantity)
		if err != nil {
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

		if _, err := s.Db.Exec(context.Background(), purchaseProduct, q, purchase.Warehouse_id, pr); err != nil {
			return fmt.Errorf("Ошибка в списании товара со склада: %w", err)
		}
	}

	return nil
}
