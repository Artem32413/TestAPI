package service

import (
	"analytics/internal/model/structs"

	"context"
)

func (s *AnalyticsService) AnalyticsTopLogic(ctx context.Context) ([]structs.TopAnalytics, error) {
	return s.repo.DisplayTop(ctx)
}
