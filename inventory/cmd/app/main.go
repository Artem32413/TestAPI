//	@title			API Go Service
//	@version		1.0
//	@description	This is a sample API Go server with Zap logging

//	@contact.name	API Support
//	@contact.url	http://www.example.com/support
//	@contact.email	support@example.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1
//	@schemes	http

package app

import (
	"inventory/cmd/host"
	"inventory/pkg/logger"

	"context"
	"os"
	"os/signal"

	_ "inventory/docs"

	"go.uber.org/zap"
)

// @Summary		Запуск приложения
// @Description	Основная точка входа для API сервиса
func main() {
	logger := logger.ZapFunc()

	if err := realMain(logger); err != nil {
		logger.Error(err.Error())
		return
	}
}

// realMain содержит основную логику приложения
//
//	@Summary		Основная логика приложения
//	@Description	Инициализирует контекст и запускает API сервер
func realMain(logger *zap.Logger) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	return host.StartMain(ctx, logger)
}
