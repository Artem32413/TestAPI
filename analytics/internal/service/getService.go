package service

import (
	"analytics/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

func (s *AnalyticsService) AnalyticsTopLogic(log *zap.Logger, ctx context.Context) ([]structs.TopAnalytics, error) {
	log.Info("Получение топ складов по выручке")
	return s.repo.DisplayTop(ctx)
}
