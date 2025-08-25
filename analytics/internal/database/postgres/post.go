package postgres

import (
	"analytics/internal/model/structs"

	"context"
	"fmt"
)

var (
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
)

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
