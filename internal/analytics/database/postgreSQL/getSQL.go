package postgreSQL

import (
	"apiGo/internal/analytics/model/structs"

	"context"
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