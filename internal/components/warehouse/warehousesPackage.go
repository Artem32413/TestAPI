package warehouse

import (
	"apiGo/internal/components"
	"net/http"
)

type Warehouses struct {
	Identifier string `json:"identifier"`
	Addr       string `json:"addr"`
}

type InventoryService struct {
	*components.Settings
}

func (s *InventoryService) AddingNewWarehouses(w http.ResponseWriter, r *http.Request) {

	var warehouses Warehouses

	if err := components.NewDec(r, &warehouses); err != nil {
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := s.Addition(warehouses); err != nil {
		s.Logger.Error(err.Error())
		return
	}

}

func (s *InventoryService) DisplayAllWarehouses(w http.ResponseWriter, r *http.Request) {

	AllWarehouses, err := s.Display()
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(AllWarehouses)
	if err != nil {
		s.Logger.Error("Ошибка в преобразовании JSON (Склады)")
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		s.Logger.Error("Ошибка в выводе данных (Склады)")
		return
	}

}
