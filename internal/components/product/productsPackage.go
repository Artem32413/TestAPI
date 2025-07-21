package product

import (
	"apiGo/internal/components"
	"net/http"

	"go.uber.org/zap"
)

type Products struct {
	Product_id  string              `json:"product_id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	KeyValue    []map[string]string `json:"keyvalue,omitempty"`
	Weight      string              `json:"weight,omitempty"`
	Barcode     string              `json:"barcode,omitempty"`
}

type InventoryService struct {
	*components.Settings
}

func (s *InventoryService) DisplayAllProducts(w http.ResponseWriter, r *http.Request) {

	prod, err := s.DisplayProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выведении всех товаров", zap.Error(err))
		return
	}

	jsData, err := components.NewMarshall(prod)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в преобразовании JSON (Товары)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выводе данных (Товары)", zap.Error(err))
		return
	}

}

func (s *InventoryService) AddingNewProducts(w http.ResponseWriter, r *http.Request) {

	var products Products

	defer r.Body.Close()

	if err := components.NewDec(r, &products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.AdditionProducts(products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *InventoryService) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var products Products

	if err := components.NewDec(r, &products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.UpdateProd(products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
