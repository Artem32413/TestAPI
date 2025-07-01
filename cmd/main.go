package main

import (
	"apiGo/config"
	"apiGo/internal/api"

	"context"
	"os"
	"os/signal"

	"go.uber.org/zap"
)

func main() {
	logger := config.ZapFunc()

	if err := realMain(logger); err != nil {
		logger.Error(err.Error())
		return
	}
}

func realMain(logger *zap.Logger) error {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	if err := api.StartMain(ctx, logger); err != nil {
		logger.Error(err.Error())
	}

	return nil
}
