package databaseConfig

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type PostgreSQL struct {
	Db *pgx.Conn
}

func ConstructorDB(ctx context.Context) (*PostgreSQL, error) {
	url := os.Getenv("DATABASE_URL")

	db, err := pgx.Connect(ctx, url)

	if err != nil {
		return nil, fmt.Errorf("Ошибка инициализации БД: %v", err)
	}

	return &PostgreSQL{
		Db: db,
	}, err
}
