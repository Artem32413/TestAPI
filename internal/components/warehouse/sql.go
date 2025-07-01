package warehouse

import "context"

var (
	addingAWarehouse = `INSERT INTO WarehousesTable (identifier, addr) VALUES ($1, $2)`
	displayWarehouse = `SELECT * FROM WarehousesTable`
)

func (s *InventoryService) Addition(warehouses Warehouses) error {

	if _, err := s.Db.Exec(context.Background(), addingAWarehouse, warehouses.Identifier, warehouses.Addr); err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) Display() ([]Warehouses, error) {

	r, err := s.Db.Query(context.Background(), displayWarehouse)
	if err != nil {
		return nil, err
	}

	var newSl []Warehouses

	for r.Next() {
		var nw Warehouses

		if err = r.Scan(&nw.Addr, &nw.Identifier); err != nil {
			return nil, err
		}

		newSl = append(newSl, Warehouses{nw.Addr, nw.Identifier})
	}

	return newSl, nil
}
