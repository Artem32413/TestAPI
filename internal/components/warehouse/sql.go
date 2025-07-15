package warehouse

import (
	"context"

	"github.com/google/uuid"
)

var (
	addingAWarehouse = `INSERT INTO WarehousesTable (warehouse_id, addr) VALUES ($1, $2)`
	displayWarehouse = `SELECT * FROM WarehousesTable`
)

func (s *InventoryService) Addition(warehouses Warehouses) error {

	if s.Db.IsClosed() {
		return s.Db.Ping(context.Background())
	}

	newUuid := uuid.New().String()

	warehouses.Warehouse_id = newUuid

	if _, err := s.Db.Exec(context.Background(), addingAWarehouse, warehouses.Warehouse_id, warehouses.Addr); err != nil {
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

		if err = r.Scan(&nw.Id, &nw.Addr, &nw.Warehouse_id); err != nil {
			return nil, err
		}

		newSl = append(newSl, Warehouses{nw.Id, nw.Addr, nw.Warehouse_id})
	}

	return newSl, nil
}
