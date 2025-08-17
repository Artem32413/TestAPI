package service

import (
	"apiGo/internal/inventory/database/postgreSQL"
	"apiGo/internal/inventory/model/structs"
	"context"
)

type InventoryService struct {
	repo *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}

func (i *InventoryService) SetPriceLogic(ctx context.Context, price structs.Inventory) error {
	return i.repo.SetPriceSQL(ctx, price)
}

func (i *InventoryService) UpdateInventoryLogic(ctx context.Context, products structs.Inventory) error {
	return i.repo.UpdateInventorySQL(ctx, products)
}

func (i *InventoryService) DiscountInventoryLogic(ctx context.Context, products structs.NewInventoryDiscount) error {
	return i.repo.DiscountInventorySQL(ctx, products)
}

func (i *InventoryService) ListOfGoodsLogic(ctx context.Context, price structs.WarehousePagination) ([]structs.ListByWarehouse, error) {
	return i.repo.ListOfGoodsLogicSQL(ctx, price)
}

func (i *InventoryService) ReceivingGoodsLogic(ctx context.Context, products structs.Inventory) (structs.AllInformationAboutTheProduct, error) {
	return i.repo.ReceivingGoodsSQL(ctx, products)
}

func (i *InventoryService) CountPriceLogic(ctx context.Context, products structs.NewInventory) (structs.SummingUp, error) {
	return i.repo.CountPriceSQL(ctx, products)
}

func (i *InventoryService) PurchaseProductLogic(ctx context.Context, products structs.NewInventory) error {
	return i.repo.PurchaseProductSQL(ctx, products)
}
