package host

import (
	"product/cmd/host/middleware"
	"product/internal/transport"

	"context"
	"net/http"

	"go.uber.org/zap"
)

func StartMain(ctx context.Context, log *zap.Logger) error {

	log.Info("Сервис Товары запущен",
		zap.Int("Порт", 8082),
	)

	mux := transport.AllHandles(ctx, log)

	s := http.Server{
		Addr:    ":8082",
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