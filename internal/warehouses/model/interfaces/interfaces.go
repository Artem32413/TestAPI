package interfaces

import (
	model "apiGo/internal/warehouses/model/structs"
	"context"
	"net/http"
)

type HandlersWarehouses interface {
	AddingNewWarehouses(w http.ResponseWriter, r *http.Request)
	DisplayAllWarehouses(w http.ResponseWriter, r *http.Request)
}

type WarehousesRepo interface {
	Addition(warehouses model.Warehouses, ctx context.Context) error
	Display(ctx context.Context) ([]model.Warehouses, error)
}