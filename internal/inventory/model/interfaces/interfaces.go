package interfaces

import (
	model "apiGo/internal/inventory/model/structs"
	"context"
	"net/http"
)

type HandlersInventory interface {
	SetPrice(w http.ResponseWriter, r *http.Request)
	UpdateInventory(w http.ResponseWriter, r *http.Request)
	DiscountInventory(w http.ResponseWriter, r *http.Request)
	ListOfGoods(w http.ResponseWriter, r *http.Request)
	ReceivingGoods(w http.ResponseWriter, r *http.Request)
	CountPrice(w http.ResponseWriter, r *http.Request)
	PurchaseProduct(w http.ResponseWriter, r *http.Request)
}

type InventoryRepo interface {
	SetPriceDB(ctx context.Context, price model.Inventory) error
	UpdateQuantity(ctx context.Context, inventory model.Inventory) error
	CreatingADiscount(ctx context.Context, discount model.NewInventoryDiscount) error
	ListProductsByWarehouse(ctx context.Context, warehouse model.WarehousePagination) ([]model.ListByWarehouse, error)
	ListProduct(ctx context.Context, product model.Inventory) (model.AllInformationAboutTheProduct, error)
	ListCount(ctx context.Context, count model.NewInventory) (model.SummingUp, error)
	Purchase(ctx context.Context, purchase model.NewInventory) error
}
