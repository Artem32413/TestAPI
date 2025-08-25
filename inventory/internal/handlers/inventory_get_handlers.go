package handlers

import (
	"inventory/internal/model/structs"
	"inventory/pkg/errors"
	"inventory/pkg/headers"
	"inventory/pkg/requests"

	"context"
	"net/http"
	"time"
)

// ListOfGoods возвращает список товаров на складе
// @Summary Список товаров
// @Description Возвращает список товаров на указанном складе с пагинацией
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body structs.WarehousePagination true "Параметры пагинации"
// @Success 200 {array} structs.ListByWarehouse "Список товаров"
// @Failure 400 "Неверные входные данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/goods/ [get]
func (i *InventoryHandler) ListOfGoods(w http.ResponseWriter, r *http.Request) {
	var product structs.WarehousePagination

	if err := requests.NewDec(r, &product); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	res, err := i.svc.ListOfGoodsLogic(i.logger, ctx, product)
	if err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	jsData, err := requests.NewMarshall(res)
	if err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(i.logger, w, jsData)
}

// ReceivingGoods возвращает полную информацию о товаре
// @Summary Информация о товаре
// @Description Возвращает полную информацию о товаре на складе
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body structs.InventorySwagger4 true "Идентификаторы склада и товара"
// @Success 200 {object} structs.AllInformationAboutTheProduct "Полная информация о товаре"
// @Failure 400 "Неверные входные данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/product/ [get]
func (i *InventoryHandler) ReceivingGoods(w http.ResponseWriter, r *http.Request) {
	var product structs.Inventory

	if err := requests.NewDec(r, &product); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	res, err := i.svc.ReceivingGoodsLogic(i.logger, ctx, product)
	if err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	jsData, err := requests.NewMarshall(res)
	if err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(i.logger, w, jsData)
}