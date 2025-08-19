package service

import (
	"apiGo/internal/warehouses/model/structs"

	"context"
)

func (w *WarehousesService) AddingNewWarehousesLogic(ctx context.Context, str structs.Warehouses) error {
	return w.repo.AddingNewWarehousesSQL(ctx, str)
}
