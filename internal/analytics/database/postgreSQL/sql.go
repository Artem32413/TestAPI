package postgreSQL

import (
	"apiGo/internal/analytics/config/databaseConfig"
	"apiGo/internal/analytics/model/structs"

	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var (
	// Аналитика
	existsWarehouse  = `SELECT EXISTS(SELECT 1 FROM Inventory WHERE warehouseId = $1)`
	productAnalytics = `SELECT 
							warehouseId,
							productId,
							SUM(sold_goods) AS total_sold,
							SUM(total_sum) AS total_revenue
						FROM 
							Analytics
						WHERE 
							warehouseId = $1
						GROUP BY 
							warehouseId, productId 
						ORDER BY 
							total_revenue DESC;`
	topWarehouses = `
						SELECT 
						w.warehouseId,
						w.addr,
						SUM(a.total_sum) AS total_revenue
					FROM 
						WarehousesTable w
					JOIN 
						Analytics a ON w.warehouseId = a.warehouseId
					GROUP BY 
						w.warehouseId, w.addr
					ORDER BY 
						total_revenue DESC
					LIMIT 10;
						`
)

type DBService struct {
    db *pgx.Conn
}

func New(db *databaseConfig.PostgreSQL) *DBService {
    return &DBService{db: db.Db}
}

func (s *DBService) DisplayAllAnalytics(ctx context.Context, str structs.Analytics) ([]structs.Analytics, error) {
	var exist bool

	if err := s.db.QueryRow(ctx, existsWarehouse, str.WarehouseId).Scan(&exist); err != nil {
		return nil, fmt.Errorf("Ошибка в проверке на существование склада")
	}

	if !exist {
		return nil, fmt.Errorf("Склад не найден")
	}

	r, err := s.db.Query(ctx, productAnalytics, str.WarehouseId)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения запроса аналитики: %w", err)
	}

	defer r.Close()

	var a []structs.Analytics

	for r.Next() {
		var NewAnalytics structs.Analytics

		if err = r.Scan(
			&NewAnalytics.WarehouseId,
			&NewAnalytics.ProductId,
			&NewAnalytics.SoldGoods,
			&NewAnalytics.TotalSum,
		); err != nil {
			return nil, fmt.Errorf("Ошибка сканирования данных аналитики: %w", err)
		}

		a = append(a, NewAnalytics)
	}

	return a, nil
}

func (s *DBService) DisplayTop(ctx context.Context) ([]structs.TopAnalytics, error) {
	r, err := s.db.Query(ctx, topWarehouses)
	if err != nil {
		return nil, err
	}

	var slAnalytic []structs.TopAnalytics

	for r.Next() {
		var a structs.TopAnalytics

		if err = r.Scan(&a.Addr, &a.WarehouseId, &a.TotalSum); err != nil {
			return nil, err
		}

		slAnalytic = append(slAnalytic, structs.TopAnalytics{Addr: a.Addr, WarehouseId: a.WarehouseId, TotalSum: a.TotalSum})
	}

	return slAnalytic, nil
}
