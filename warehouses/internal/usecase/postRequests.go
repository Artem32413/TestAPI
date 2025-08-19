package usecase

import (
	model "warehouses/internal/model/structs"
	"warehouses/pkg/errors"
	"warehouses/pkg/headers"
	"warehouses/pkg/requests"

	"context"
	"net/http"
	"time"
)

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

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := wr.svc.AddingNewWarehousesLogic(ctx, warehouses); err != nil {
		errors.HandleError(wr.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(wr.logger, w, []byte("Успешное добавление склада"))
}
