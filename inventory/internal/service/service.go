package service

import (
	"inventory/internal/database/postgreSQL"
)

type InventoryService struct {
	repo *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}











