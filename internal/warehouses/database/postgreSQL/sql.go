package postgreSQL

import (
	"apiGo/internal/warehouses/config/settings"
	"apiGo/internal/warehouses/model/structs"
	"context"

	"github.com/google/uuid"
)

var (
	addingAWarehouse = `INSERT INTO WarehousesTable (warehouse_id, addr) VALUES ($1, $2)`
	displayWarehouse = `SELECT * FROM WarehousesTable`
)

type InventoryService struct {
	*settings.Settings
}


func (s *InventoryService) Addition(ctx context.Context, warehouses model.Warehouses) error {

	if s.Db.IsClosed() {
		return s.Db.Ping(ctx)
	}

	newUuid := uuid.New().String()

	warehouses.Warehouse_id = newUuid

	if _, err := s.Db.Exec(ctx, addingAWarehouse, warehouses.Warehouse_id, warehouses.Addr); err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) Display(ctx context.Context) ([]model.Warehouses, error) {

	r, err := s.Db.Query(ctx, displayWarehouse)
	if err != nil {
		return nil, err
	}

	var newSl []model.Warehouses

	for r.Next() {
		var nw model.Warehouses

		if err = r.Scan(&nw.Id, &nw.Addr, &nw.Warehouse_id); err != nil {
			return nil, err
		}

		newSl = append(newSl, model.Warehouses{Id: nw.Id, Addr: nw.Addr, Warehouse_id: nw.Warehouse_id})
	}

	return newSl, nil
}
