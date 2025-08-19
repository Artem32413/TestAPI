package zapConfig

import (
	"inventory/pkg/logger"

	"fmt"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type ZapConf struct {
	Logger *zap.Logger
}

func SetZap() (*ZapConf, error) {
	logger := logger.ZapFunc()

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Ошибка в инициализации zap logger: %v", err)
	}

	return &ZapConf{
		Logger: logger,
	}, nil
}
