package analytics

import (
	"apiGo/internal/components"
	"net/http"
)

type Analytics struct {
	Warehouse_id string  `json:"warehouse_id"`
	Product_id   string  `json:"product_id"`
	SoldGoods    int     `json:"soldgoods"`
	TotalSum     float64 `json:"totalsum"`
}

type TopAnalytics struct {
	Addr         string  `json:"addr"`
	Warehouse_id string  `json:"warehouse_id"`
	TotalSum     float64 `json:"totalsum"`
}

type InventoryService struct {
	*components.Settings
}

func (s *InventoryService) AnalyticsAll(w http.ResponseWriter, r *http.Request) {
	var a Analytics

	if err := components.NewDec(r, &a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	res, err := s.DisplayAllAnalytics(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в преобразовании JSON (Склады)")
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выводе данных (Склады)")
		return
	}
}

func (s *InventoryService) Top(w http.ResponseWriter, r *http.Request) {
	res, err := s.DisplayTop()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в преобразовании JSON (Склады)")
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выводе данных (Склады)")
		return
	}
}
