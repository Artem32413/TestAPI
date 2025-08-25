package service

import (
	"warehouses/internal/database/postgres"
)

type WarehousesService struct {
	repo *postgres.DBService
}

func NewWarehousesService(repo *postgres.DBService) *WarehousesService {
	return &WarehousesService{
		repo: repo,
	}
}
