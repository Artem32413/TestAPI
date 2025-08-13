package service

import (
	"apiGo/internal/warehouses/database/postgreSQL"
	"apiGo/internal/warehouses/model/structs"

	"context"
)

type WarehousesService struct {
    repo   *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *WarehousesService {
    return &WarehousesService{
        repo: repo,
    }
}

func (w *WarehousesService) AddingNewWarehousesLogic(ctx context.Context, str structs.Warehouses)  error {
	return w.repo.AddingNewWarehousesSQL(ctx, str)
}

func (w *WarehousesService) DisplayAllWarehousesLogic(ctx context.Context) ([]structs.Warehouses, error) {
	return w.repo.DisplayAllWarehousesSQL(ctx)
}