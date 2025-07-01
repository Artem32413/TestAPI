package analytics

import (
	"apiGo/internal/components"
	"net/http"
)

type Analytics struct {
	Warehouse_id string  `json:"warehouses"`
	Product_id   string  `json:"products"`
	SoldGoods    int     `json:"soldgoods"`
	TotalSum     float64 `json:"totalsum"`
}

type InventoryService struct {
    *components.Settings
}

func (s *InventoryService) AnalyticsAll(w http.ResponseWriter, r *http.Request) {
	res, err := s.DisplayAllAnalytics()
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
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

func (s *InventoryService) Top(w http.ResponseWriter, r *http.Request) {
	res, err := s.DisplayTop()
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
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
