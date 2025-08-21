package host

import (
	"warehouses/cmd/host/middleware"
	"warehouses/internal/transport"

	"context"
	"net/http"

	"go.uber.org/zap"
)

func StartMain(ctx context.Context, log *zap.Logger) error {

	log.Info("Сервис Склады запущен",
		zap.Int("Порт", 8081),
	)

	mux := transport.AllHandles(ctx, log)

	s := http.Server{
		Addr:    ":8081",
		Handler: middleware.LoggingMiddleware(mux),
	}

	go func() {
		<-ctx.Done()
		log.Info("Сервер завершен")
		s.Shutdown(ctx)
	}()

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
