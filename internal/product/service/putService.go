package service

import (
	"apiGo/internal/product/model/structs"

	"context"
)

func (p *ProductService) UpdateProductLogic(ctx context.Context, products structs.Products) error {
	return p.repo.UpdateProductSQL(ctx, products)
}
