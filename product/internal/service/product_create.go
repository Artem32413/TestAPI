package service

import (
	"product/internal/model/structs"

	"context"

	"go.uber.org/zap"
)

func (p *ProductService) AddingNewProductsLogic(log *zap.Logger, ctx context.Context, products structs.Products) error {
	log.Info("Добавление нового товара")
	return p.repo.AddingNewProductsSQL(ctx, products)
}
