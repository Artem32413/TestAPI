package handlers

import (
	"analytics/internal/service"

	"go.uber.org/zap"
)

type AnalyticsHandler struct {
	svc    *service.AnalyticsService
	logger *zap.Logger
}

func NewAnalyticsHandler(svc *service.AnalyticsService, logger *zap.Logger) *AnalyticsHandler {
	return &AnalyticsHandler{
		svc:    svc,
		logger: logger,
	}
}
