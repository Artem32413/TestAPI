package service

import (
	"inventory/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

func (i *InventoryService) ListOfGoodsLogic(log *zap.Logger, ctx context.Context, price structs.WarehousePagination) ([]structs.ListByWarehouse, error) {
	log.Info("Получение списка товаров по конкретному складу с пагинацией")
	return i.repo.ListOfGoodsLogicSQL(log, ctx, price)
}

func (i *InventoryService) ReceivingGoodsLogic(log *zap.Logger, ctx context.Context, products structs.Inventory) (structs.AllInformationAboutTheProduct, error) {
	log.Info("Получение товара на складе")
	return i.repo.ReceivingGoodsSQL(log, ctx, products)
}