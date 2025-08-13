package service

import (
	"apiGo/internal/analytics/database/postgreSQL"
	"apiGo/internal/analytics/model/structs"

	"context"
)

type AnalyticsService struct {
    repo   *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *AnalyticsService {
    return &AnalyticsService{
        repo: repo,
    }
}

func (s *AnalyticsService) AnalyticsAllLogic(ctx context.Context, str structs.Analytics) ([]structs.Analytics, error) {
	return s.repo.DisplayAllAnalytics(ctx, str)
}

func (s *AnalyticsService) AnalyticsTopLogic(ctx context.Context) ([]structs.TopAnalytics, error) {
	return s.repo.DisplayTop(ctx)
}