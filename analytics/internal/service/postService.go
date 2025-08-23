package service

import (
	"analytics/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

func (s *AnalyticsService) AnalyticsAllLogic(log *zap.Logger, ctx context.Context, str structs.Analytics) ([]structs.Analytics, error) {
	log.Info("Получение общей аналитики")
	return s.repo.DisplayAllAnalytics(ctx, str)
}
