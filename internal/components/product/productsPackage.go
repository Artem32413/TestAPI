package product

import (
	"apiGo/internal/components"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type Products struct {
	Identifier  string          `json:"identifier"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	KeyValue    map[uint]string `json:"keyvalue"`
	Weight      string          `json:"weight"`
	Barcode     string          `json:"barcode"`
}

type InventoryService struct {
    *components.Settings
}


func (s *InventoryService) DisplayAllProducts(w http.ResponseWriter, r *http.Request) {
	
	prod, err := s.DisplayProducts()
	if err != nil {
		s.Logger.Error("Ошибка в выведении всех товаров")
		return
	}

	jsData, err := components.NewMarshall(prod)

	if err != nil {
		s.Logger.Error("Ошибка в преобразовании JSON (Товары)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		s.Logger.Error("Ошибка в выводе данных (Товары)", zap.Error(err))
		return
	}

}

func (s *InventoryService) AddingNewProducts(w http.ResponseWriter, r *http.Request) {

	var products Products

	defer r.Body.Close()

	if err := components.NewDec(r, &products); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	if err := s.AdditionProducts(products); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *InventoryService) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	identifier, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || identifier < 1 {
		http.NotFound(w, r)
		s.Logger.Error("Такого identifier не найдено")
		return
	}

	var products Products

	defer r.Body.Close()

	if err := components.NewDec(r, &products); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	if err := s.UpdateProd(products, identifier); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
