package postgres

import (
	"warehouses/internal/config/database"

	"github.com/jackc/pgx/v5"
)

type DBService struct {
	db *pgx.Conn
}

func NewWarehouseRepository(db *database.PostgreSQL) *DBService {
	return &DBService{db: db.Conn}
}
