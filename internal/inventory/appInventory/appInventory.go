package appInventory

import (
	"apiGo/internal/inventory/model/structs"
	"apiGo/internal/inventory/service"
	"apiGo/pkg/errors"
	"apiGo/pkg/headers"
	"apiGo/pkg/requests"
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type InventoryHandler struct {
	svc    *service.InventoryService
	logger *zap.Logger
}

func New(svc *service.InventoryService, logger *zap.Logger) *InventoryHandler {
	return &InventoryHandler{
		svc:    svc,
		logger: logger,
	}
}

// SetPrice устанавливает цену товара на складе
// @Summary Установить цену товара
// @Description Устанавливает цену для конкретного товара на указанном складе
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body InventorySwagger1 true "Данные установки цены"
// @Success 200 "Цена успешно установлена"
// @Failure 400 "Неверные входные данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/price/ [patch]
func (i *InventoryHandler) SetPrice(w http.ResponseWriter, r *http.Request) {

	var price structs.Inventory

	if err := requests.NewDec(r, &price); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := i.svc.SetPriceLogic(ctx, price); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(i.logger, w, []byte("Цена успешно установлена"))
}

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

	if err := i.svc.UpdateInventoryLogic(ctx, inventory); err != nil {
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

	if err := i.svc.DiscountInventoryLogic(ctx, discount); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(i.logger, w, []byte("Скидка успешно применена"))
}

// ListOfGoods возвращает список товаров на складе
// @Summary Список товаров
// @Description Возвращает список товаров на указанном складе с пагинацией
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body WarehousePagination true "Параметры пагинации"
// @Success 200 {array} ListByWarehouse "Список товаров"
// @Failure 400 "Неверные входные данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/goods/ [patch]
func (i *InventoryHandler) ListOfGoods(w http.ResponseWriter, r *http.Request) {
	var product structs.WarehousePagination

	if err := requests.NewDec(r, &product); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := i.svc.ListOfGoodsLogic(ctx, product)
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
// @Param input body InventorySwagger4 true "Идентификаторы склада и товара"
// @Success 200 {object} AllInformationAboutTheProduct "Полная информация о товаре"
// @Failure 400 "Неверные входные данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/product/ [patch]
func (i *InventoryHandler) ReceivingGoods(w http.ResponseWriter, r *http.Request) {
	var product structs.Inventory

	if err := requests.NewDec(r, &product); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := i.svc.ReceivingGoodsLogic(ctx, product)
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

	res, err := i.svc.CountPriceLogic(ctx, count)
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

// PurchaseProduct обрабатывает покупку товаров
// @Summary Покупка товаров
// @Description Обрабатывает покупку товаров и уменьшает их количество на складе
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body NewInventory true "Данные о покупке"
// @Success 200 "Покупка успешно обработана"
// @Failure 400 "Неверные входные данные или недостаточно товара"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/purchase/ [patch]
func (i *InventoryHandler) PurchaseProduct(w http.ResponseWriter, r *http.Request) {
	var purchase structs.NewInventory

	if err := requests.NewDec(r, &purchase); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := i.svc.PurchaseProductLogic(ctx, purchase); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(i.logger, w, []byte("Покупка успешно обработана"))
}
