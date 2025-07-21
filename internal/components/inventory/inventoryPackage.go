package inventory

import (
	"apiGo/internal/components"
	"net/http"

	"go.uber.org/zap"
)

type WarehousePagination struct {
	Warehouse_id string `json:"warehouse_id"`
	Limit        int    `json:"limit"`  //Количество элементов в блоке
	Offset       int    `json:"offset"` //Номер блока
}

type Inventory struct {
	Warehouse_id string  `json:"warehouse_id"`
	Product_id   string  `json:"product_id"`
	Quantity     int     `json:"quantity,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Discount     float64 `json:"discount,omitempty"`
}

type NewInventory struct {
	Warehouse_id string              `json:"warehouse_id"`
	Product      []ProductInventory2 `json:"product"`
}

type NewInventoryDiscount struct {
	Warehouse_id string   `json:"warehouse_id"`
	Product_id   []string `json:"product_id"`
	Discount     float64  `json:"discount"`
}

type ProductInventory2 struct {
	Product_id string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

type InventoryService struct {
	*components.Settings
}

func (s *InventoryService) SetPrice(w http.ResponseWriter, r *http.Request) {

	var price Inventory

	if err := components.NewDec(r, &price); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.SetPriceDB(price); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *InventoryService) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	var inventory Inventory

	if err := components.NewDec(r, &inventory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.UpdateQuantity(inventory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (s *InventoryService) DiscountInventory(w http.ResponseWriter, r *http.Request) {

	var discount NewInventoryDiscount

	if err := components.NewDec(r, &discount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.CreatingADiscount(discount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *InventoryService) ListOfGoods(w http.ResponseWriter, r *http.Request) {
	var product WarehousePagination

	if err := components.NewDec(r, &product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	res, err := s.ListProductsByWarehouse(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в преобразовании JSON (Товары со склада)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выводе данных (Товары со склада)", zap.Error(err))
		return
	}
}

func (s *InventoryService) ReceivingGoods(w http.ResponseWriter, r *http.Request) {
	var product Inventory

	if err := components.NewDec(r, &product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	res, err := s.ListProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в преобразовании JSON (Товара со склада)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выводе данных (Товара со склада)", zap.Error(err))
		return
	}
}

func (s *InventoryService) CountPrice(w http.ResponseWriter, r *http.Request) {
	var count NewInventory

	if err := components.NewDec(r, &count); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	res, err := s.ListCount(count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в преобразовании JSON (Подсчёта)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выводе данных (Подсчёта)", zap.Error(err))
		return
	}
}

func (s *InventoryService) PurchaseProduct(w http.ResponseWriter, r *http.Request) {
	var purchase NewInventory

	if err := components.NewDec(r, &purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.Purchase(purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
