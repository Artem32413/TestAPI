package service

import (
	"inventory/internal/database/postgres"
)

type InventoryService struct {
	repo *postgres.DBService
}

func NewInventoryService(repo *postgres.DBService) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}











