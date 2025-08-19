package service

import (
	"apiGo/internal/warehouses/database/postgreSQL"
)

type WarehousesService struct {
	repo *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *WarehousesService {
	return &WarehousesService{
		repo: repo,
	}
}
