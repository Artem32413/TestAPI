package api

import (
	"context"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func StartMain(ctx context.Context, logger *zap.Logger) error {

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	logger.Info("Сервер запущен")

	mux := AllHandles()

	s := http.Server{
		Addr:    addr,
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

		requestID := r.Header.Get("x-request-id")

		if requestID == "" {
			requestID = uuid.New().String()
		}

		logger := logrus.WithField("request_id", requestID)

		ctx := context.WithValue(r.Context(), keyRequestID{}, requestID)
		ctx = context.WithValue(ctx, "logger", logger)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
