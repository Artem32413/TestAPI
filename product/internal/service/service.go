package service

import (
	"product/internal/database/postgres"
)

type ProductService struct {
	repo *postgres.DBService
}

func NewProductService(repo *postgres.DBService) *ProductService {
	return &ProductService{
		repo: repo,
	}
}
