package postgres

import (
	"analytics/internal/model/structs"

	"context"
)

var (
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
					LIMIT 10;`
)

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
