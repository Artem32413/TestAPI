package appInventory

import (
	"apiGo/internal/inventory/model/structs"
	"apiGo/pkg/errors"
	"apiGo/pkg/headers"
	"apiGo/pkg/requests"

	"context"
	"net/http"
	"time"
)

// UpdateInventory обновляет количество товара
// @Summary Обновить количество товара
// @Description Увеличивает количество товара на складе
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body InventorySwagger2 true "Данные обновления количества"
// @Success 200 "Количество успешно обновлено"
// @Failure 400 "Неверные входные данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/updateQuantity/ [patch]
func (i *InventoryHandler) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	var inventory structs.Inventory

	if err := requests.NewDec(r, &inventory); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := i.svc.UpdateInventoryLogic(i.logger, ctx, inventory); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(i.logger, w, []byte("Количество успешно обновлено"))
}

// DiscountInventory устанавливает скидку на товары
// @Summary Установить скидку
// @Description Устанавливает скидку на указанные товары на складе
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body NewInventoryDiscount true "Данные для скидки"
// @Success 200 "Скидка успешно применена"
// @Failure 400 "Неверные входные данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/discount/ [patch]
func (i *InventoryHandler) DiscountInventory(w http.ResponseWriter, r *http.Request) {

	var discount structs.NewInventoryDiscount

	if err := requests.NewDec(r, &discount); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := i.svc.DiscountInventoryLogic(i.logger, ctx, discount); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(i.logger, w, []byte("Скидка успешно применена"))
}

// CountPrice подсчитывает стоимость товаров
// @Summary Подсчет стоимости
// @Description Возвращает сумму стоимости указанных товаров с учетом скидок
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body NewInventory true "Список товаров для подсчета"
// @Success 200 {object} InventorySwagger5 "Итоговая стоимость"
// @Failure 400 "Неверно введены данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/count/ [patch]
func (i *InventoryHandler) CountPrice(w http.ResponseWriter, r *http.Request) {
	var count structs.NewInventory

	if err := requests.NewDec(r, &count); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := i.svc.CountPriceLogic(i.logger, ctx, count)
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

