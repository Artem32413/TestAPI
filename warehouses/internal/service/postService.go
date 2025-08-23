package service

import (
	"warehouses/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

func (w *WarehousesService) AddingNewWarehousesLogic(log *zap.Logger, ctx context.Context, str structs.Warehouses) error {
	log.Info("Добавление нового склада")
	return w.repo.AddingNewWarehousesSQL(ctx, str)
}
