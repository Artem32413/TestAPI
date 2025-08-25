package service

import (
	"inventory/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

func (i *InventoryService) PurchaseProductLogic(log *zap.Logger, ctx context.Context, products structs.NewInventory) error {
	log.Info("Покупка товаров со склада")
	return i.repo.PurchaseProductSQL(log, ctx, products)
}
