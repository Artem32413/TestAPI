package host

import (
	"context"
	"fmt"
	"inventory/cmd/host/closer"
	"inventory/cmd/host/middleware"
	"inventory/internal/transport"
	"net/http"
	"time"

	"go.uber.org/zap"
)

const (
	port            = ":8083"
	shutdownTimeout = 5 * time.Second
)

func StartMain(ctx context.Context, log *zap.Logger) error {
	var (
		mux = transport.AllHandles(ctx, log)
		srv = &http.Server{
			Addr:    port,
			Handler: middleware.LoggingMiddleware(mux),
		}
		c = &closer.Closer{}
	)

	c.Add(srv.Shutdown)

	c.Add(func(ctx context.Context) error {
		time.Sleep(3 * time.Second)

		return nil
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("сервер закрыт: ",
				zap.Error(err),
			)
		}
	}()

	log.Info("сервер Инвентаризация слушает ",
		zap.String("порт", port),
	)
	<-ctx.Done()

	log.Info("корректное завершение работы сервера")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := c.Close(shutdownCtx); err != nil {
		return fmt.Errorf("закрытие: %v", err)
	}

	return nil
}
