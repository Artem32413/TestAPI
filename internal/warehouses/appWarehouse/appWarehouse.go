package appWarehouse

import (
	"apiGo/internal/components"
	"apiGo/internal/warehouses/config/settings"
	"apiGo/internal/warehouses/model/interfaces"
	model "apiGo/internal/warehouses/model/structs"
	"context"
	"fmt"
	"net/http"
	"time"
)

type InventoryService struct {
	*settings.Settings

}

type WarehousesSwagger struct {
	Addr string `json:"addr" example:"fghverv4446"`
}

var st struct {
	interfaces.WarehousesRepo
	components.InventoryService
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
	var warehouses model.Warehouses

	if err := components.NewDec(r, &warehouses); err != nil {
		st.HandleError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := st.Addition(warehouses, ctx); err != nil {
		st.HandleError(w, err, http.StatusBadRequest)
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
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	AllWarehouses, err := st.Display(ctx)
	if err != nil {
		st.HandleError(w, err, http.StatusBadRequest)
		return
	}

	jsData, err := components.NewMarshall(AllWarehouses)
	if err != nil {
		st.HandleError(w, fmt.Errorf("Ошибка в преобразовании JSON (Склады)"), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		st.HandleError(w, fmt.Errorf("Ошибка в выводе данных (Склады)"), http.StatusBadRequest)
		return
	}
}
