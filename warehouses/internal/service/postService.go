package service

import (
	"warehouses/internal/model/structs"

	"context"
)

func (w *WarehousesService) AddingNewWarehousesLogic(ctx context.Context, str structs.Warehouses) error {
	return w.repo.AddingNewWarehousesSQL(ctx, str)
}
