package service

import (
	"inventory/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

func (i *InventoryService) UpdateInventoryLogic(log *zap.Logger, ctx context.Context, products structs.Inventory) error {
	log.Info("Обновление количества определённого товара на складе")
	return i.repo.UpdateInventorySQL(log, ctx, products)
}

func (i *InventoryService) DiscountInventoryLogic(log *zap.Logger, ctx context.Context, products structs.NewInventoryDiscount) error {
	log.Info("Создание скидки на определённый список товаров на складе")
	return i.repo.DiscountInventorySQL(log, ctx, products)
}

func (i *InventoryService) CountPriceLogic(log *zap.Logger, ctx context.Context, products structs.NewInventory) (structs.SummingUp, error) {
	log.Info("Получения подсчёта")
	return i.repo.CountPriceSQL(log, ctx, products)
}
