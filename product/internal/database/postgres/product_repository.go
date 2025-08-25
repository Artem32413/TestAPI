package postgres

import (
	"product/internal/config/database"

	"github.com/jackc/pgx/v5"
)

type DBService struct {
	db *pgx.Conn
}

func NewProductRepository(db *database.PostgreSQL) *DBService {
	return &DBService{db: db.Conn}
}
