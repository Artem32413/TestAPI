package warehouse

import (
	"apiGo/internal/components"
	"net/http"
)

// Warehouses представляет склад в системе
// swagger:model Warehouse
type Warehouses struct {
	// Уникальный идентификатор склада
	Id int `json:"id" example:"1"`

	// Внешний идентификатор склада
	Warehouse_id string `json:"warehouse_id" example:"WH-001"`

	// Физический адрес склада
	Addr string `json:"addr" example:"ул. Складская, д.1"`
}

type InventoryService struct {
	*components.Settings
}


type WarehousesSwagger struct {
	Addr string `json:"addr" example:"fghverv4446"`
}
// AddingNewWarehouses добавляет новый склад в систему
// @Summary Добавить новый склад
// @Description Добавляет новый склад в систему складирования
// @Tags Склады
// @Accept  json
// @Produce  json
// @Param warehouse body WarehousesSwagger true "Данные склада"
// @Success 200 "Склад успешно добавлен"
// @Failure 400 "Некорректные входные данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /warehouses/add/ [post]
func (s *InventoryService) AddingNewWarehouses(w http.ResponseWriter, r *http.Request) {
	var warehouses Warehouses

	if err := components.NewDec(r, &warehouses); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := s.Addition(warehouses); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}
}

// DisplayAllWarehouses получает список всех складов
// @Summary Получить все склады
// @Description Возвращает список всех складов в системе
// @Tags Склады
// @Produce  json
// @Success 200 {array} Warehouses "Список складов"
// @Failure 400 "Ошибка запроса"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /warehouses/all/ [get]
func (s *InventoryService) DisplayAllWarehouses(w http.ResponseWriter, r *http.Request) {
	AllWarehouses, err := s.Display()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(AllWarehouses)
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
