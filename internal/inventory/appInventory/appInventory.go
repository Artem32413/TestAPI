package appInventory

import (
	"apiGo/internal/components"
	"apiGo/internal/inventory/config/settings"
	"apiGo/internal/inventory/model/interfaces"
	model "apiGo/internal/inventory/model/structs"
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type InventoryService struct {
	*settings.Settings
}

var st struct{
	interfaces.InventoryRepo
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
func (s *InventoryService) SetPrice(w http.ResponseWriter, r *http.Request) {

	var price model.Inventory

	if err := components.NewDec(r, &price); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := st.SetPriceDB(ctx, price); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
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
func (s *InventoryService) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	var inventory model.Inventory

	if err := components.NewDec(r, &inventory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := st.UpdateQuantity(ctx, inventory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

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
func (s *InventoryService) DiscountInventory(w http.ResponseWriter, r *http.Request) {

	var discount model.NewInventoryDiscount

	if err := components.NewDec(r, &discount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := st.CreatingADiscount(ctx, discount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
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
func (s *InventoryService) ListOfGoods(w http.ResponseWriter, r *http.Request) {
	var product model.WarehousePagination

	if err := components.NewDec(r, &product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := st.ListProductsByWarehouse(ctx, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в преобразовании JSON (Товары со склада)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выводе данных (Товары со склада)", zap.Error(err))
		return
	}
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
func (s *InventoryService) ReceivingGoods(w http.ResponseWriter, r *http.Request) {
	var product model.Inventory

	if err := components.NewDec(r, &product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := st.ListProduct(ctx, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в преобразовании JSON (Товара со склада)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выводе данных (Товара со склада)", zap.Error(err))
		return
	}
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
func (s *InventoryService) CountPrice(w http.ResponseWriter, r *http.Request) {
	var count model.NewInventory

	if err := components.NewDec(r, &count); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := st.ListCount(ctx, count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	jsData, err := components.NewMarshall(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в преобразовании JSON (Подсчёта)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выводе данных (Подсчёта)", zap.Error(err))
		return
	}
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
func (s *InventoryService) PurchaseProduct(w http.ResponseWriter, r *http.Request) {
	var purchase model.NewInventory

	if err := components.NewDec(r, &purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := st.Purchase(ctx, purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
