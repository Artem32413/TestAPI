package service

import (
	"apiGo/internal/analytics/model/structs"

	"context"
)

func (s *AnalyticsService) AnalyticsTopLogic(ctx context.Context) ([]structs.TopAnalytics, error) {
	return s.repo.DisplayTop(ctx)
}
