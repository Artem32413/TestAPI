package service

import (
	"warehouses/internal/database/postgreSQL"
)

type WarehousesService struct {
	repo *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *WarehousesService {
	return &WarehousesService{
		repo: repo,
	}
}
