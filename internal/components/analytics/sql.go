package analytics

import "context"

var (
	// Аналитика
	productAnalytics = `
SELECT 
    p.product_id,
    p.name,
    p.barcode,
    SUM(a.sold_goods) AS total_sold,
    SUM(a.total_sum) AS total_sum,
    w.identifier AS warehouse_identifier,
    w.addr AS warehouse_address
FROM 
    analytics a
JOIN 
    products p ON a.product_id = p.product_id
JOIN 
    warehouses w ON a.warehouse_id = w.identifier
GROUP BY 
    p.product_id, p.name, p.barcode, w.identifier, w.addr
`

	topWarehouses = `
SELECT 
    w.identifier,
    w.addr,
    SUM(a.total_sum) AS total_revenue
FROM 
    warehouses w
JOIN 
    analytics a ON w.identifier = a.warehouse_id
GROUP BY 
    w.identifier, w.addr
ORDER BY 
    total_revenue DESC
LIMIT 10
`
)

func (s *InventoryService) DisplayAllAnalytics() ([]Analytics, error) {
	r, err := s.Db.Query(context.Background(), productAnalytics)
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
	r, err := s.Db.Query(context.Background(), topWarehouses)
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
