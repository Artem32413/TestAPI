package service

import (
	"analytics/internal/database/postgres"
)

type AnalyticsService struct {
	repo *postgres.DBService
}

func NewAnalyticsService(repo *postgres.DBService) *AnalyticsService {
	return &AnalyticsService{
		repo: repo,
	}
}
