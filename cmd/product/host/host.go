package host

import (
	"apiGo/cmd/product/host/middleware"
	"apiGo/internal/product/transport"

	"context"
	"net/http"

	"go.uber.org/zap"
)

func StartMain(ctx context.Context, logger *zap.Logger) error {

	logger.Info("Сервер запущен")

	mux := transport.AllHandles(ctx, logger)

	s := http.Server{
		Addr:    ":8082",
		Handler: middleware.LoggingMiddleware(mux),
	}

	go func() {
		<-ctx.Done()
		logger.Info("Сервер завершен")
		s.Shutdown(ctx)
	}()

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}