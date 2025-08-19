package service

import (
	"apiGo/internal/inventory/model/structs"

	"context"

	"go.uber.org/zap"
)

func (i *InventoryService) SetPriceLogic(log *zap.Logger, ctx context.Context, price structs.Inventory) error {
	log.Info("Создание связи товара и склада")
	return i.repo.SetPriceSQL(log, ctx, price)
}