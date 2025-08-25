package database 

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type PostgreSQL struct {
	Conn *pgx.Conn
}

func NewPostgreSQL(ctx context.Context) (*PostgreSQL, error) {
	connStr := os.Getenv("DATABASE_URL")

	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("Ошибка инициализации БД: %v", err)
	}

	return &PostgreSQL{
		Conn: db,
	}, err
}
