package service

import (
	"apiGo/internal/analytics/database/postgreSQL"
)

type AnalyticsService struct {
	repo *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *AnalyticsService {
	return &AnalyticsService{
		repo: repo,
	}
}
