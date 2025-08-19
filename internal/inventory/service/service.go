package service

import (
	"apiGo/internal/inventory/database/postgreSQL"
)

type InventoryService struct {
	repo *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}











