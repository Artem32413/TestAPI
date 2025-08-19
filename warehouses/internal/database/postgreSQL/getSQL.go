package postgreSQL

import (
	"warehouses/internal/model/structs"

	"context"
)

func (d *DBService) DisplayAllWarehousesSQL(ctx context.Context) ([]structs.Warehouses, error) {

	r, err := d.db.Query(ctx, displayWarehouse)
	if err != nil {
		return nil, err
	}

	var newSl []structs.Warehouses

	for r.Next() {
		var nw structs.Warehouses

		if err = r.Scan(&nw.Id, &nw.Addr, &nw.WarehouseId); err != nil {
			return nil, err
		}

		newSl = append(newSl, structs.Warehouses{Id: nw.Id, Addr: nw.Addr, WarehouseId: nw.WarehouseId})
	}

	return newSl, nil
}
