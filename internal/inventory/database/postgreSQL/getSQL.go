package postgreSQL

import (
	"apiGo/internal/inventory/database/postgreSQL/helpFunc"
	"apiGo/internal/inventory/model/structs"
	"fmt"

	"context"

	"go.uber.org/zap"
)

func (d *DBService) ListOfGoodsLogicSQL(log *zap.Logger, ctx context.Context, warehouse structs.WarehousePagination) ([]structs.ListByWarehouse, error) {

	if exist := helpFunc.Exists(log, d.db, ctx, 5, structs.Inventory{}, structs.NewInventoryDiscount{}, structs.NewInventory{}, warehouse); !exist {
		return nil, fmt.Errorf("Склад с ID %s не найден", warehouse.WarehouseId)
	}

	rows, err := d.db.Query(
		context.Background(),
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

	if exist := helpFunc.Exists(log, d.db, ctx, 1, product, structs.NewInventoryDiscount{}, structs.NewInventory{}, structs.WarehousePagination{}); !exist {
		return structs.AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка в проверке на существование склада и товара")
	}

	var n structs.AllInformationAboutTheProduct

	if err := d.db.QueryRow(ctx, listInventory, product.WarehouseId, product.ProductId).Scan(&n.Price, &n.Discount, &n.Quantity); err != nil {
		return structs.AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка с запросом остатка %w", err)
	}

	if err := d.db.QueryRow(ctx, oneProduct, product.ProductId).Scan(&n.ProductId, &n.Name, &n.Description, &n.Weight, &n.Barcode); err != nil {
		return structs.AllInformationAboutTheProduct{}, fmt.Errorf("Ошибка с запросом товара %w", err)
	}

	attrs, err := helpFunc.GetAllAttributes(d.db, []string{n.ProductId})
	if err != nil {
		return n, fmt.Errorf("Ошибка получения характеристик: %w", err)
	}

	if productAttrs, exists := attrs[n.ProductId]; exists {
		n.Characteristics = helpFunc.ConvertAttributesToSlice(productAttrs)
	}

	return n, nil
}
