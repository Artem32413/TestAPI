package postgres

import (
	"inventory/internal/config/database"

	"github.com/jackc/pgx/v5"
)

type DBService struct {
	db *pgx.Conn
}

func NewInventoryRepository(db *database.PostgreSQL) *DBService {
	return &DBService{db: db.Conn}
}
