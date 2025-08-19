package service

import (
	"analytics/internal/model/structs"

	"context"
)

func (s *AnalyticsService) AnalyticsAllLogic(ctx context.Context, str structs.Analytics) ([]structs.Analytics, error) {
	return s.repo.DisplayAllAnalytics(ctx, str)
}
