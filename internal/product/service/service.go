package service

import (
	"apiGo/internal/product/database/postgreSQL"
	"apiGo/internal/product/model/structs"

	"context"
)

type ProductService struct {
	repo *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (p *ProductService) DisplayAllProductsLogic(ctx context.Context) ([]structs.Products, error) {
	return p.repo.DisplayAllProductsSQL(ctx)
}

func (p *ProductService) AddingNewProductsLogic(ctx context.Context, products structs.Products) error {
	return p.repo.AddingNewProductsSQL(ctx, products)
}

func (p *ProductService) UpdateProductLogic(ctx context.Context, products structs.Products) error {
	return p.repo.UpdateProductSQL(ctx, products)
}
