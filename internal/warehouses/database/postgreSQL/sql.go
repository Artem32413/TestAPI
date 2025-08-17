package postgreSQL

import (
	"apiGo/internal/warehouses/config/databaseConfig"
	"apiGo/internal/warehouses/model/structs"

	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	addingAWarehouse = `INSERT INTO Warehouses (warehouseId, addr) VALUES ($1, $2)`
	displayWarehouse = `SELECT * FROM Warehouses`
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

	warehouses.WarehouseId = newUuid

	if _, err := d.db.Exec(ctx, addingAWarehouse, warehouses.WarehouseId, warehouses.Addr); err != nil {
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

		if err = r.Scan(&nw.Id, &nw.Addr, &nw.WarehouseId); err != nil {
			return nil, err
		}

		newSl = append(newSl, structs.Warehouses{Id: nw.Id, Addr: nw.Addr, WarehouseId: nw.WarehouseId})
	}

	return newSl, nil
}
