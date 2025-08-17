package postgreSQL

import (
	"apiGo/internal/inventory/config/databaseConfig"
	"apiGo/internal/inventory/database/postgreSQL/helpFunc"
	"apiGo/internal/inventory/model/structs"

	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type DBService struct {
	db *pgx.Conn
}

func New(db *databaseConfig.PostgreSQL) *DBService {
	return &DBService{db: db.Db}
}

func (d *DBService) SetPriceSQL(ctx context.Context, price structs.Inventory) error {

	if exist := helpFunc.Exists(d.db, ctx, 1, price, structs.NewInventoryDiscount{}, structs.NewInventory{}, structs.WarehousePagination{}); exist {
		if _, err := d.db.Exec(ctx, priceUpdate, price.Price, price.WarehouseId, price.ProductId); err != nil {
			return err
		}
	} else {
		if _, err := d.db.Exec(ctx, priceInsert, price.WarehouseId, price.ProductId, 0, price.Price, 0); err != nil {
			return err
		}
	}

	return nil
}

func (d *DBService) UpdateInventorySQL(ctx context.Context, inventory structs.Inventory) error {

	if exist := helpFunc.Exists(d.db, ctx, 1, inventory, structs.NewInventoryDiscount{}, structs.NewInventory{}, structs.WarehousePagination{}); !exist {
		return fmt.Errorf("Ошибка в проверке на существование склада и товара")
	}

	if _, err := d.db.Exec(ctx, updateQuantity, inventory.Quantity, inventory.WarehouseId, inventory.ProductId); err != nil {
		return err
	}

	return nil
}

func (d *DBService) DiscountInventorySQL(ctx context.Context, discount structs.NewInventoryDiscount) error {

	if exist := helpFunc.Exists(d.db, ctx, 3, structs.Inventory{}, discount, structs.NewInventory{}, structs.WarehousePagination{}); !exist {
		return fmt.Errorf("Склад не найден")
	}

	for _, el := range discount.ProductId {
		if _, err := d.db.Exec(ctx, creatingADiscount, discount.Discount, discount.WarehouseId, el); err != nil {
			return err
		}
	}

	return nil
}

func (d *DBService) ListOfGoodsLogicSQL(ctx context.Context, warehouse structs.WarehousePagination) ([]structs.ListByWarehouse, error) {

	if exist := helpFunc.Exists(d.db, ctx, 5, structs.Inventory{}, structs.NewInventoryDiscount{}, structs.NewInventory{}, warehouse); !exist {
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

func (d *DBService) ReceivingGoodsSQL(ctx context.Context, product structs.Inventory) (structs.AllInformationAboutTheProduct, error) {

	if exist := helpFunc.Exists(d.db, ctx, 1, product, structs.NewInventoryDiscount{}, structs.NewInventory{}, structs.WarehousePagination{}); !exist {
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

func (d *DBService) CountPriceSQL(ctx context.Context, count structs.NewInventory) (structs.SummingUp, error) {
	var result structs.SummingUp

	if exist := helpFunc.Exists(d.db, ctx, 4, structs.Inventory{}, structs.NewInventoryDiscount{}, count, structs.WarehousePagination{}); !exist {
		return result, fmt.Errorf("Склад с ID %s не найден", count.WarehouseId)
	}

	if len(count.Product) == 0 {
		return result, fmt.Errorf("Список товаров не может быть пустым")
	}

	productIDs, quantities := helpFunc.Slices(count)

	if err := d.db.QueryRow(
		ctx,
		listCount,
		productIDs,
		quantities,
		count.WarehouseId,
	).Scan(&result.Sum); err != nil {
		return result, fmt.Errorf("Ошибка расчета суммы: %w", err)
	}

	return result, nil
}

func (d *DBService) PurchaseProductSQL(ctx context.Context, purchase structs.NewInventory) error {

	if exist := helpFunc.Exists(d.db, ctx, 4, structs.Inventory{}, structs.NewInventoryDiscount{}, purchase, structs.WarehousePagination{}); !exist {
		return fmt.Errorf("Склад не найден")
	}

	productIDs, quantities := helpFunc.Slices(purchase)

	var quantity int

	for _, productId := range productIDs {
		if err := d.db.QueryRow(context.Background(), quantityCheck, purchase.WarehouseId, productId).Scan(&quantity); err != nil {
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

		if _, err := d.db.Exec(context.Background(), purchaseProduct, q, purchase.WarehouseId, pr); err != nil {
			return fmt.Errorf("Ошибка в списании товара со склада: %w", err)
		}

		if _, err := d.db.Exec(context.Background(), purchaseProductAnalytics, purchase.WarehouseId, pr, q); err != nil {
			return fmt.Errorf("Ошибка в записи данных в аналитику: %w", err)
		}
	}

	return nil
}
