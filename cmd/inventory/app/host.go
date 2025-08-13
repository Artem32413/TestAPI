package app

import (
	"apiGo/internal/inventory/transport"
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func StartMain(ctx context.Context, logger *zap.Logger) error {

	logger.Info("Сервер запущен")

	mux := transport.AllHandles()

	s := http.Server{
		Addr:    ":8083",
		Handler: LoggingMiddleware(mux),
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

type keyRequestID struct{}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		newUuid := r.Header.Get("x-request-id")

		newUuid = uuid.New().String()

		logger := logrus.WithField("request_id", newUuid)

		ctx := context.WithValue(r.Context(), keyRequestID{}, newUuid)
		ctx = context.WithValue(ctx, "logger", logger)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
