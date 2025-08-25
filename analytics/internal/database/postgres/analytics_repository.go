package postgres

import (
	"analytics/internal/config/database"

	"github.com/jackc/pgx/v5"
)

type DBService struct {
	db *pgx.Conn
}

func NewAnalyticsRepository(db *database.PostgreSQL) *DBService {
	return &DBService{db: db.Conn}
}
