package analytics

import (
	"apiGo/internal/components"
	"net/http"
)

// Analytics представляет данные аналитики по продажам
// @Description Структура данных аналитики по продажам товаров на складе
type Analytics struct {
	Warehouse_id string  `json:"warehouse_id" example:"WH-001"` // Идентификатор склада
	Product_id   string  `json:"product_id" example:"PRD-1001"` // Идентификатор товара
	SoldGoods    int     `json:"soldgoods" example:"50"`        // Количество проданных товаров
	TotalSum     float64 `json:"totalsum" example:"12500.50"`   // Общая сумма продаж
}

// TopAnalytics представляет данные топовых складов по выручке
// @Description Структура данных топовых складов по выручке
type TopAnalytics struct {
	Addr         string  `json:"addr" example:"ул. Ленина, 10"` // Адрес склада
	Warehouse_id string  `json:"warehouse_id" example:"WH-001"` // Идентификатор склада
	TotalSum     float64 `json:"totalsum" example:"150000.75"`  // Общая выручка склада
}

type InventoryService struct {
	*components.Settings
}

type AnalyticsSwagger struct {
	Warehouse_id string `json:"warehouse_id" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"` // Идентификатор склада
}

// AnalyticsAll обрабатывает запрос аналитики продаж
// @Summary Получить аналитику продаж
// @Description Возвращает аналитику продаж по указанным параметрам
// @Tags Analytics
// @Accept json
// @Produce json
// @Param request body AnalyticsSwagger true "Параметры запроса аналитики"
// @Success 200 {array} Analytics "Успешный ответ"
// @Failure 400 "Ошибка в запросе"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /analytics/ [post]
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

// Top обрабатывает запрос топовых складов
// @Summary Получить топ складов по выручке
// @Description Возвращает список складов с наибольшей выручкой
// @Tags Analytics
// @Accept json
// @Produce json
// @Success 200 {array} TopAnalytics "Список топовых складов"
// @Failure 400 "Ошибка в запросе"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /analytics/top/ [get]
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