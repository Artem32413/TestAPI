package postgreSQL

import (
	"apiGo/internal/warehouses/config/databaseConfig"
	"apiGo/internal/warehouses/model/structs"

	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	addingAWarehouse = `INSERT INTO WarehousesTable (warehouse_id, addr) VALUES ($1, $2)`
	displayWarehouse = `SELECT * FROM WarehousesTable`
)

type DBService struct {
	db *pgx.Conn
}

func New(db *databaseConfig.PostgreSQL) *DBService {
	return &DBService{db: db.Db}
}

func (d *DBService) AddingNewWarehousesSQL(ctx context.Context, warehouses structs.Warehouses) error {

	if d.db.IsClosed() {
		return d.db.Ping(ctx)
	}

	newUuid := uuid.New().String()

	warehouses.Warehouse_id = newUuid

	if _, err := d.db.Exec(ctx, addingAWarehouse, warehouses.Warehouse_id, warehouses.Addr); err != nil {
		return err
	}

	return nil
}

func (d *DBService) DisplayAllWarehousesSQL(ctx context.Context) ([]structs.Warehouses, error) {

	r, err := d.db.Query(ctx, displayWarehouse)
	if err != nil {
		return nil, err
	}

	var newSl []structs.Warehouses

	for r.Next() {
		var nw structs.Warehouses

		if err = r.Scan(&nw.Id, &nw.Addr, &nw.Warehouse_id); err != nil {
			return nil, err
		}

		newSl = append(newSl, structs.Warehouses{Id: nw.Id, Addr: nw.Addr, Warehouse_id: nw.Warehouse_id})
	}

	return newSl, nil
}
