package interfaces

import (
	model "apiGo/internal/product/model/structs"
	"context"
	"net/http"
)

type HandlersProducts interface {
	DisplayAllProducts(w http.ResponseWriter, r *http.Request)
	AddingNewProducts(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}

type ProductsRepo interface {
	DisplayProducts(ctx context.Context) ([]model.Products, error)
	AdditionProducts(ctx context.Context, products model.Products) error
	UpdateProd(ctx context.Context, products model.Products) error
}
