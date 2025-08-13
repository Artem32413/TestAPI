package appWarehouse

import (
	model "apiGo/internal/warehouses/model/structs"
	"apiGo/internal/warehouses/service"
	"apiGo/pkg/errors"
	"apiGo/pkg/requests"
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type WarehousesHandler struct {
	svc    *service.WarehousesService
	logger *zap.Logger
}

func New(svc *service.WarehousesService, logger *zap.Logger) *WarehousesHandler {
	return &WarehousesHandler{
		svc:    svc,
		logger: logger,
	}
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
func (wr *WarehousesHandler) AddingNewWarehouses(w http.ResponseWriter, r *http.Request) {
	var warehouses model.Warehouses

	if err := requests.NewDec(r, &warehouses); err != nil {
		errors.HandleError(wr.logger, w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := wr.svc.AddingNewWarehousesLogic(ctx, warehouses); err != nil {
		errors.HandleError(wr.logger, w, err, http.StatusBadRequest)
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
func (wr *WarehousesHandler) DisplayAllWarehouses(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	AllWarehouses, err := wr.svc.DisplayAllWarehousesLogic(ctx)
	if err != nil {
		errors.HandleError(wr.logger, w, err, http.StatusBadRequest)
		return
	}

	jsData, err := requests.NewMarshall(AllWarehouses)
	if err != nil {
		errors.HandleError(wr.logger, w, fmt.Errorf("Ошибка в преобразовании JSON (Склады)"), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		errors.HandleError(wr.logger, w, fmt.Errorf("Ошибка в выводе данных (Склады)"), http.StatusBadRequest)
		return
	}
}