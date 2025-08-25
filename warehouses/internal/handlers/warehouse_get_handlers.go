package handlers

import (
	"warehouses/pkg/errors"
	"warehouses/pkg/headers"
	"warehouses/pkg/requests"

	"context"
	"fmt"
	"net/http"
	"time"
)

// DisplayAllWarehouses получает список всех складов
// @Summary Получить все склады
// @Description Возвращает список всех складов в системе
// @Tags Склады
// @Produce  json
// @Success 200 {array} structs.Warehouses "Список складов"
// @Failure 400 "Ошибка запроса"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /warehouses/all/ [get]
func (wr *WarehousesHandler) DisplayAllWarehouses(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	AllWarehouses, err := wr.svc.DisplayAllWarehousesLogic(wr.logger, ctx)
	if err != nil {
		errors.HandleError(wr.logger, w, err, http.StatusBadRequest)
		return
	}

	jsData, err := requests.NewMarshall(AllWarehouses)
	if err != nil {
		errors.HandleError(wr.logger, w, fmt.Errorf("Ошибка в преобразовании JSON (Склады)"), http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(wr.logger, w, jsData)
}
