package service

import (
	"apiGo/internal/product/model/structs"

	"context"
)

func (p *ProductService) AddingNewProductsLogic(ctx context.Context, products structs.Products) error {
	return p.repo.AddingNewProductsSQL(ctx, products)
}
