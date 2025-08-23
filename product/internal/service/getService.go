package service

import (
	"product/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

func (p *ProductService) DisplayAllProductsLogic(log *zap.Logger, ctx context.Context) ([]structs.Products, error) {
	log.Info("Получение списка всех товаров")
	return p.repo.DisplayAllProductsSQL(ctx)
}
