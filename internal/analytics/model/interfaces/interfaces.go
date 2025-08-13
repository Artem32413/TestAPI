package interfaces

import (
	model "apiGo/internal/analytics/model/structs"
	"context"
	"net/http"
)

type HandlersAnalytics interface {
	AnalyticsAll(w http.ResponseWriter, r *http.Request)
	AnalyticsTop(w http.ResponseWriter, r *http.Request)
}

type AnalyticsRepo interface {
	DisplayAllAnalytics(ctx context.Context, str model.Analytics) ([]model.Analytics, error)
	DisplayTop(ctx context.Context) ([]model.TopAnalytics, error)
}