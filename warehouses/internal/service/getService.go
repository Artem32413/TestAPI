package service

import (
	"warehouses/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

func (w *WarehousesService) DisplayAllWarehousesLogic(log *zap.Logger, ctx context.Context) ([]structs.Warehouses, error) {
	log.Info("Получение списка всех складов")
	return w.repo.DisplayAllWarehousesSQL(ctx)
}
