package service

import (
	"apiGo/internal/product/model/structs"

	"context"
)

func (p *ProductService) DisplayAllProductsLogic(ctx context.Context) ([]structs.Products, error) {
	return p.repo.DisplayAllProductsSQL(ctx)
}
