package postgres

import (
	"fmt"
	"inventory/internal/database/postgres/attributes"
	"inventory/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

var (
	listOfGoods = `SELECT 
						p.productId,
						p.name,
						i.price,
						ROUND(i.price * (1 - i.discount/100), 2) as discounted_price
						FROM inventory i
						JOIN products p ON i.productId = p.productId
						WHERE i.warehouseId = $1
						LIMIT $2 OFFSET $3`
	listInventory = `	SELECT 
						COALESCE(quantity, 0) AS quantity,
						COALESCE(price, 0) AS price,
						COALESCE(discount, 0) AS discount 
						FROM Inventory 
						WHERE warehouseId = $1 AND productId = $2`
	oneProduct = `SELECT 
						p.productId,
						p.name,
						p.description,
						p.weight,
						p.barcode
						FROM 
						Products p
						WHERE productId = $1`
)

func (d *DBService) ListOfGoodsLogicSQL(log *zap.Logger, ctx context.Context, warehouse structs.WarehousePagination) ([]structs.ListByWarehouse, error) {

	if exist := attributes.Exists(log, d.db, ctx, 5, structs.Inventory{}, structs.NewInventoryDiscount{}, structs.NewInventory{}, warehouse); !exist {
		return nil, fmt.Errorf("Склад с ID %s не найден", warehouse.WarehouseId)
	}

	rows, err := d.db.Query(
		ctx,
		listOfGoods,
		warehouse.WarehouseId,
		warehouse.Limit,
		warehouse.Offset,
	)

	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения запроса: %w", err)
	}

	defer rows.Close()

	var products []structs.ListByWarehouse

	for rows.Next() {
		var p structs.ListByWarehouse

		err := rows.Scan(
			&p.ProductId,
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

func (d *DBService) ReceivingGoodsSQL(log *zap.Logger, ctx context.Context, product structs.Inventory) (structs.AllInformationAboutTheProduct, error) {

	if exist := attributes.Exists(log, d.db, ctx, 1, product, structs.NewInventoryDiscount{}, structs.NewInventory{}, structs.WarehousePagination{}); !exist {
		return structs.AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка в проверке на существование склада и товара")
	}

	var n structs.AllInformationAboutTheProduct

	if err := d.db.QueryRow(ctx, listInventory, product.WarehouseId, product.ProductId).Scan(&n.Price, &n.Discount, &n.Quantity); err != nil {
		return structs.AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка с запросом остатка %w", err)
	}

	if err := d.db.QueryRow(ctx, oneProduct, product.ProductId).Scan(&n.ProductId, &n.Name, &n.Description, &n.Weight, &n.Barcode); err != nil {
		return structs.AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка с запросом товара %w", err)
	}

	attrs, err := attributes.GetAllAttributes(d.db, []string{n.ProductId})
	if err != nil {
		return n, fmt.Errorf("Ошибка получения характеристик: %w", err)
	}

	if productAttrs, exists := attrs[n.ProductId]; exists {
		n.Characteristics = attributes.ConvertAttributesToSlice(productAttrs)
	}

	return n, nil
}
