package postgreSQL

import (
	"apiGo/internal/warehouses/config/databaseConfig"

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
