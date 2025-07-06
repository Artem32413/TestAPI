package warehouse

import (
	"context"

	"github.com/google/uuid"
)

var (
	addingAWarehouse = `INSERT INTO WarehousesTable (identifier, addr) VALUES ($1, $2)`
	displayWarehouse = `SELECT * FROM WarehousesTable`
)

func (s *InventoryService) Addition(warehouses Warehouses) error {

	if s.Db.IsClosed() {
		return s.Db.Ping(context.Background())
	}

	newUuid := uuid.New().String()

	warehouses.Identifier = newUuid

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

		if err = r.Scan(&nw.Id, &nw.Addr, &nw.Identifier); err != nil {
			return nil, err
		}

		newSl = append(newSl, Warehouses{nw.Id, nw.Addr, nw.Identifier})
	}

	return newSl, nil
}
