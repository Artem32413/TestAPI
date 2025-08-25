package service

import (
	"product/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

func (p *ProductService) UpdateProductLogic(log *zap.Logger, ctx context.Context, products structs.Products) error {
	log.Info("Обновление данных товара")
	return p.repo.UpdateProductSQL(ctx, products)
}
