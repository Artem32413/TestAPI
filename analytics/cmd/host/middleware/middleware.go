package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

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
