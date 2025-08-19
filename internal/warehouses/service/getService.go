package service

import (
	"apiGo/internal/warehouses/model/structs"

	"context"
)

func (w *WarehousesService) DisplayAllWarehousesLogic(ctx context.Context) ([]structs.Warehouses, error) {
	return w.repo.DisplayAllWarehousesSQL(ctx)
}
