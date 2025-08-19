package service

import (
	"apiGo/internal/product/database/postgreSQL"
)

type ProductService struct {
	repo *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *ProductService {
	return &ProductService{
		repo: repo,
	}
}
