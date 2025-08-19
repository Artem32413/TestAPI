package service

import (
	"product/internal/model/structs"

	"context"
)

func (p *ProductService) DisplayAllProductsLogic(ctx context.Context) ([]structs.Products, error) {
	return p.repo.DisplayAllProductsSQL(ctx)
}
