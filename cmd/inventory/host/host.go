package host

import (
	"apiGo/cmd/inventory/host/middleware"
	"apiGo/internal/inventory/transport"
	"context"
	"net/http"

	"go.uber.org/zap"
)

func StartMain(ctx context.Context, log *zap.Logger) error {

	log.Info("Сервис Инвентаризация запущен",
		zap.Int("Порт", 8083),
	)

	mux := transport.AllHandles(ctx, log)

	s := http.Server{
		Addr:    ":8083",
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
