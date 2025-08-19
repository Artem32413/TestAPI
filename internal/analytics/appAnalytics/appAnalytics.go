package appAnalytics

import (
	"apiGo/internal/analytics/service"

	"go.uber.org/zap"
)

type AnalyticsHandler struct {
	svc    *service.AnalyticsService
	logger *zap.Logger
}

func New(svc *service.AnalyticsService, logger *zap.Logger) *AnalyticsHandler {
	return &AnalyticsHandler{
		svc:    svc,
		logger: logger,
	}
}
