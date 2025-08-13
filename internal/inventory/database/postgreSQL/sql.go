package postgreSQL

import (
	"apiGo/internal/inventory/config/settings"
	"apiGo/internal/inventory/database/postgreSQL/split"
	model "apiGo/internal/inventory/model/structs"
	"context"
	"fmt"
)

var (
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
	listInventory = `	SELECT 
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
	purchaseProductAnalytics = `
								INSERT INTO analytics (warehouse_id, product_id, sold_goods, total_sum)
								SELECT 
									$1::text, 
									$2::text, 
									$3::integer,
									$3::integer * (
										SELECT price * (1 - COALESCE(discount, 0)) 
										FROM Inventory 
										WHERE product_id = $2::text
									)
								ON CONFLICT (warehouse_id, product_id) 
								DO UPDATE SET
									sold_goods = analytics.sold_goods + EXCLUDED.sold_goods,
									total_sum = analytics.total_sum + EXCLUDED.total_sum`
)

var spl struct{
	split.SplitMethods
}

type InventoryService struct {
	*settings.Settings
}

func (s *InventoryService) SetPriceDB(price model.Inventory) error {

	if exist := spl.Exists(1, price, model.NewInventoryDiscount{}, model.NewInventory{}, model.WarehousePagination{}); exist {
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

func (s *InventoryService) UpdateQuantity(inventory model.Inventory) error {

	if exist := spl.Exists(1, inventory, model.NewInventoryDiscount{}, model.NewInventory{}, model.WarehousePagination{}); !exist {
		return fmt.Errorf("Ошибка в проверке на существование склада и товара")
	}

	if _, err := s.Db.Exec(context.Background(), updateQuantity, inventory.Quantity, inventory.Warehouse_id, inventory.Product_id); err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) CreatingADiscount(discount model.NewInventoryDiscount) error {

	if exist := spl.Exists(3, model.Inventory{}, discount, model.NewInventory{}, model.WarehousePagination{}); !exist {
		return fmt.Errorf("Склад не найден")
	}

	for _, el := range discount.Product_id {
		if _, err := s.Db.Exec(context.Background(), creatingADiscount, discount.Discount, discount.Warehouse_id, el); err != nil {
			return err
		}
	}

	return nil
}

func (s *InventoryService) ListProductsByWarehouse(warehouse model.WarehousePagination) ([]model.ListByWarehouse, error) {

	if exist := spl.Exists(5, model.Inventory{}, model.NewInventoryDiscount{}, model.NewInventory{}, warehouse); !exist {
		return nil, fmt.Errorf("Склад с ID %s не найден", warehouse.Warehouse_id)
	}

	rows, err := s.Db.Query(
		context.Background(),
		listOfGoods,
		warehouse.Warehouse_id,
		warehouse.Limit,
		warehouse.Offset,
	)

	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения запроса: %w", err)
	}

	defer rows.Close()

	var products []model.ListByWarehouse

	for rows.Next() {
		var p model.ListByWarehouse

		err := rows.Scan(
			&p.Product_id,
			&p.Name,
			&p.Price,
			&p.Discount,
		)

		if err != nil {
			return nil, fmt.Errorf("Ошибка сканирования данных: %w", err)
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка при обработке результатов: %w", err)
	}

	return products, nil
}

func (s *InventoryService) ListProduct(product model.Inventory) (model.AllInformationAboutTheProduct, error) {

	if exist := spl.Exists(1, product, model.NewInventoryDiscount{}, model.NewInventory{}, model.WarehousePagination{}); !exist {
		return model.AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка в проверке на существование склада и товара")
	}

	var n model.AllInformationAboutTheProduct

	if err := s.Db.QueryRow(context.Background(), listInventory, product.Warehouse_id, product.Product_id).Scan(&n.Price, &n.Discount, &n.Quantity); err != nil {
		return model.AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка с запросом остатка %w", err)
	}

	if err := s.Db.QueryRow(context.Background(), oneProduct, product.Product_id).Scan(&n.Product_id, &n.Name, &n.Description, &n.Weight, &n.Barcode); err != nil {
		return model.AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка с запросом товара %w", err)
	}

	attrs, err := spl.GetAllAttributes([]string{n.Product_id})
	if err != nil {
		return n, fmt.Errorf("Ошибка получения характеристик: %w", err)
	}

	if productAttrs, exists := attrs[n.Product_id]; exists {
		n.Characteristics = split.ConvertAttributesToSlice(productAttrs)
	}

	return n, nil
}

func (s *InventoryService) ListCount(count model.NewInventory) (model.SummingUp, error) {
	var result model.SummingUp

	if exist := spl.Exists(4, model.Inventory{}, model.NewInventoryDiscount{}, count, model.WarehousePagination{}); !exist {
		return result, fmt.Errorf("Склад с ID %s не найден", count.Warehouse_id)
	}

	if len(count.Product) == 0 {
		return result, fmt.Errorf("Список товаров не может быть пустым")
	}

	productIDs, quantities := split.Slices(count)

	if err := s.Db.QueryRow(
		context.Background(),
		listCount,
		productIDs,
		quantities,
		count.Warehouse_id,
	).Scan(&result.Sum); err != nil {
		return result, fmt.Errorf("Ошибка расчета суммы: %w", err)
	}

	return result, nil
}

func (s *InventoryService) Purchase(purchase model.NewInventory) error {

	if exist := spl.Exists(4, model.Inventory{}, model.NewInventoryDiscount{}, purchase, model.WarehousePagination{}); !exist {
		return fmt.Errorf("Склад не найден")
	}

	productIDs, quantities := split.Slices(purchase)

	var quantity int

	for _, product_id := range productIDs {
		if err := s.Db.QueryRow(context.Background(), quantityCheck, purchase.Warehouse_id, product_id).Scan(&quantity); err != nil {
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

		if _, err := s.Db.Exec(context.Background(), purchaseProductAnalytics, purchase.Warehouse_id, pr, q); err != nil {
			return fmt.Errorf("Ошибка в записи данных в аналитику: %w", err)
		}
	}

	return nil
}
