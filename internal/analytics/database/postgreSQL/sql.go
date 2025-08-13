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
	existsWarehouse  = `SELECT EXISTS(SELECT 1 FROM Inventory WHERE warehouse_id = $1)`
	productAnalytics = `SELECT 
							warehouse_id,
							product_id,
							SUM(sold_goods) AS total_sold,
							SUM(total_sum) AS total_revenue
						FROM 
							Analytics
						WHERE 
							warehouse_id = $1
						GROUP BY 
							warehouse_id, product_id 
						ORDER BY 
							total_revenue DESC;`
	topWarehouses = `
						SELECT 
						w.warehouse_id,
						w.addr,
						SUM(a.total_sum) AS total_revenue
					FROM 
						WarehousesTable w
					JOIN 
						Analytics a ON w.warehouse_id = a.warehouse_id
					GROUP BY 
						w.warehouse_id, w.addr
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

func (s *DBService) DisplayAllAnalytics(ctx context.Context, str model.Analytics) ([]model.Analytics, error) {
	var exist bool

	if err := s.db.QueryRow(ctx, existsWarehouse, str.Warehouse_id).Scan(&exist); err != nil {
		return nil, fmt.Errorf("Ошибка в проверке на существование склада")
	}

	if !exist {
		return nil, fmt.Errorf("Склад не найден")
	}

	r, err := s.db.Query(ctx, productAnalytics, str.Warehouse_id)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения запроса аналитики: %w", err)
	}

	defer r.Close()

	var a []model.Analytics

	for r.Next() {
		var NewAnalytics model.Analytics

		if err = r.Scan(
			&NewAnalytics.Warehouse_id,
			&NewAnalytics.Product_id,
			&NewAnalytics.SoldGoods,
			&NewAnalytics.TotalSum,
		); err != nil {
			return nil, fmt.Errorf("Ошибка сканирования данных аналитики: %w", err)
		}

		a = append(a, NewAnalytics)
	}

	return a, nil
}

func (s *DBService) DisplayTop(ctx context.Context) ([]model.TopAnalytics, error) {
	r, err := s.db.Query(ctx, topWarehouses)
	if err != nil {
		return nil, err
	}

	var slAnalytic []model.TopAnalytics

	for r.Next() {
		var a model.TopAnalytics

		if err = r.Scan(&a.Addr, &a.Warehouse_id, &a.TotalSum); err != nil {
			return nil, err
		}

		slAnalytic = append(slAnalytic, model.TopAnalytics{Addr: a.Addr, Warehouse_id: a.Warehouse_id, TotalSum: a.TotalSum})
	}

	return slAnalytic, nil
}
