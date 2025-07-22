package analytics

import (
	"context"
	"fmt"
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

func (s *InventoryService) DisplayAllAnalytics(str Analytics) ([]Analytics, error) {
	var exist bool

	if err := s.Db.QueryRow(context.Background(), existsWarehouse, str.Warehouse_id).Scan(&exist); err != nil {
		s.Logger.Error("Ошибка в проверке на существование склада")
	}

	if !exist {
		return nil, fmt.Errorf("Склад не найден")
	}

	r, err := s.Db.Query(context.Background(), productAnalytics, str.Warehouse_id)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса аналитики: %w", err)
	}

	defer r.Close()

	var a []Analytics

	for r.Next() {
		var NewAnalytics Analytics

		err = r.Scan(
			&NewAnalytics.Warehouse_id,
			&NewAnalytics.Product_id,
			&NewAnalytics.SoldGoods,
			&NewAnalytics.TotalSum,
		)

		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования данных аналитики: %w", err)
		}

		a = append(a, NewAnalytics)
	}

	return a, nil
}

func (s *InventoryService) DisplayTop() ([]TopAnalytics, error) {
	r, err := s.Db.Query(context.Background(), topWarehouses)
	if err != nil {
		return nil, err
	}

	var slAnalytic []TopAnalytics

	for r.Next() {
		var a TopAnalytics
		if err = r.Scan(&a.Addr, &a.Warehouse_id, &a.TotalSum); err != nil {
			return nil, err
		}

		slAnalytic = append(slAnalytic, TopAnalytics{a.Addr, a.Warehouse_id, a.TotalSum})
	}

	return slAnalytic, nil
}
