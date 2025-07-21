package warehouse

import (
	"apiGo/internal/components"
	"net/http"
)

type Warehouses struct {
	Id           int    `json:"id"`
	Warehouse_id string `json:"warehouse_id"`
	Addr         string `json:"addr"`
}

type InventoryService struct {
	*components.Settings
}

func (s *InventoryService) AddingNewWarehouses(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var warehouses Warehouses

		if err := components.NewDec(r, &warehouses); err != nil {
			s.Logger.Error(err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)

		if err := s.Addition(warehouses); err != nil {
			s.Logger.Error(err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s *InventoryService) DisplayAllWarehouses(w http.ResponseWriter, r *http.Request) {

	AllWarehouses, err := s.Display()
	if err != nil {
		s.Logger.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsData, err := components.NewMarshall(AllWarehouses)
	if err != nil {
		s.Logger.Error("Ошибка в преобразовании JSON (Склады)")
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		s.Logger.Error("Ошибка в выводе данных (Склады)")
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
