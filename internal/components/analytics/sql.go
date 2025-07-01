package analytics

import "context"

var (
	// Аналитика
	analytics = `SELECT * FROM Analytics`
	analyticsTop = `...`
)

func (s *InventoryService) DisplayAllAnalytics() ([]Analytics, error) {
	r, err := s.Db.Query(context.Background(), analytics)
	if err != nil {
		return nil, err
	}

	var slAnalytic []Analytics

	for r.Next() {
		var a Analytics
		if err = r.Scan(&a.Warehouse_id, &a.Product_id, &a.SoldGoods, &a.TotalSum); err != nil {
			return nil, err
		}

		slAnalytic = append(slAnalytic, Analytics{a.Warehouse_id, a.Product_id, a.SoldGoods, a.TotalSum})
	}

	return slAnalytic, nil
}

func (s *InventoryService) DisplayTop() ([]Analytics, error) {
	r, err := s.Db.Query(context.Background(), analyticsTop)
	if err != nil {
		return nil, err
	}

	var slAnalytic []Analytics

	for r.Next() {
		var a Analytics
		if err = r.Scan(&a.Warehouse_id, &a.Product_id, &a.SoldGoods, &a.TotalSum); err != nil {
			return nil, err
		}

		slAnalytic = append(slAnalytic, Analytics{a.Warehouse_id, a.Product_id, a.SoldGoods, a.TotalSum})
	}

	return slAnalytic, nil
}