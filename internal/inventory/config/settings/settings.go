package settings

import (
	"apiGo/internal/analytics/config/zapConfig"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Settings struct {
	Db     *pgx.Conn
	Logger *zap.Logger
}

func Set() (*Settings, error) {

	logger := zapConfig.ZapFunc()

	if err := godotenv.Load(); err != nil {
		logger.Error(err.Error())
	}

	url := os.Getenv("DATABASE_URL")

	logger.Info(url)

	db, err := pgx.Connect(context.Background(), url)

	if err != nil {
		logger.Error("Ошибка в соединении с PostgreSQL")
		return &Settings{
			Logger: logger,
		}, err
	}

	return &Settings{
		Db:     db,
		Logger: logger,
	}, nil
}
