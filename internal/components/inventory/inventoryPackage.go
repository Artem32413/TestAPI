package inventory

import (
	"apiGo/internal/components"
	"net/http"

	"go.uber.org/zap"
)

type Inventory struct {
	Warehouse_id string  `json:"warehouses"`
	Product_id   string  `json:"products"`
	Quantity     int     `json:"quantity,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Discount     float64 `json:"discount,omitempty"`
}

type NewInventory struct {
	Warehouse_id string `json:"warehouses"`
	Product_id   []int  `json:"products"`
	Quantity     int    `json:"quantity"`
}

type InventoryService struct {
    *components.Settings
}


func (s *InventoryService) Connection(w http.ResponseWriter, r *http.Request) {

	var price Inventory

	defer r.Body.Close()

	if err := components.NewDec(r, &price); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	if err := s.ConnectingTable(price); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *InventoryService) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	var inventory Inventory

	if err := components.NewDec(r, &inventory); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	if err := s.UpdateQuantity(inventory); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (s *InventoryService) DiscountInventory(w http.ResponseWriter, r *http.Request) {

	var discount Inventory

	if err := components.NewDec(r, &discount); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	if err := s.CreatingADiscount(discount); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *InventoryService) ListOfGoods(w http.ResponseWriter, r *http.Request) {
	var product Inventory

	if err := components.NewDec(r, &product); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	res, err := s.List(product, 1, 2)
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		s.Logger.Error("Ошибка в преобразовании JSON (Товары со склада)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		s.Logger.Error("Ошибка в выводе данных (Товары со склада)", zap.Error(err))
		return
	}
}

func (s *InventoryService) ReceivingGoods(w http.ResponseWriter, r *http.Request) {
	var product Inventory

	if err := components.NewDec(r, &product); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	res, err := s.ListProduct(product)
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		s.Logger.Error("Ошибка в преобразовании JSON (Товара со склада)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		s.Logger.Error("Ошибка в выводе данных (Товара со склада)", zap.Error(err))
		return
	}
}

func (s *InventoryService) CountPrice(w http.ResponseWriter, r *http.Request) {
	var count Inventory

	if err := components.NewDec(r, &count); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	res, err := s.ListCount(count)
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		s.Logger.Error("Ошибка в преобразовании JSON (Подсчёта)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		s.Logger.Error("Ошибка в выводе данных (Подсчёта)", zap.Error(err))
		return
	}
}

func (s *InventoryService) PurchaseProduct(w http.ResponseWriter, r *http.Request) {
	var purchase NewInventory

	if err := components.NewDec(r, &purchase); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	if err := s.Purchase(purchase); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
