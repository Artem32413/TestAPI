package inventory

import (
	"apiGo/internal/components"
	"net/http"

	"go.uber.org/zap"
)

// WarehousePagination структура пагинации склада
// @Description Параметры пагинации для списка товаров на складе
type WarehousePagination struct {
	Warehouse_id string `json:"warehouse_id" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	Limit        int    `json:"limit" example:"10"` // Количество элементов на странице
	Offset       int    `json:"offset" example:"0"` // Смещение (номер страницы * limit)
}

// Inventory структура данных товара на складе
// @Description Основная информация о товаре на складе
type Inventory struct {
	Warehouse_id string  `json:"warehouse_id" example:"WH-001"`
	Product_id   string  `json:"product_id" example:"PRD-1001"`
	Quantity     int     `json:"quantity,omitempty"`
	Price        float64 `json:"price,omitempty" example:"999.99"`
	Discount     float64 `json:"discount,omitempty"`
}

// NewInventory структура для создания/обновления инвентаря
// @Description Данные для создания или обновления инвентарной записи
type NewInventory struct {
	Warehouse_id string              `json:"warehouse_id" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	Product      []ProductInventory2 `json:"product"`
}

// NewInventoryDiscount структура для установки скидки
// @Description Данные для установки скидки на товары
type NewInventoryDiscount struct {
	Warehouse_id string   `json:"warehouse_id" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	Product_id   []string `json:"product_id" example:"881670ed-8195-43d6-b678-3741bd3289ae,25f8f0ba-75f0-4caa-8648-27f6be3254b3"`
	Discount     float64  `json:"discount" example:"0.2"` // Размер скидки (0.2 = 20%)
}

// ProductInventory2 структура товара для операций
// @Description Упрощенная структура товара для операций
type ProductInventory2 struct {
	Product_id string `json:"product_id" example:"881670ed-8195-43d6-b678-3741bd3289ae"`
	Quantity   int    `json:"quantity" example:"5"`
}

type InventoryService struct {
	*components.Settings
}

type InventorySwagger1 struct {
	Warehouse_id string  `json:"warehouse_id" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	Product_id   string  `json:"product_id" example:"36c6ecf5-1e13-4cef-a591-bc8368fa0e60"`
	Price        float64 `json:"price,omitempty" example:"999.99"`
}

type InventorySwagger2 struct {
	Warehouse_id string `json:"warehouse_id" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	Product_id   string `json:"product_id" example:"36c6ecf5-1e13-4cef-a591-bc8368fa0e60"`
	Quantity     int    `json:"quantity" example:"5"`
}

type InventorySwagger3 struct {
	Warehouse_id string   `json:"warehouse_id" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	Product_id   []string `json:"product_id" example:"36c6ecf5-1e13-4cef-a591-bc8368fa0e60"`
	Quantity     int      `json:"quantity" example:"5"`
}

type InventorySwagger4 struct {
	Warehouse_id string `json:"warehouse_id" example:"9be99fa0-2cd1-4c4e-863d-5115095fbd09"`
	Product_id   string `json:"product_id" example:"36c6ecf5-1e13-4cef-a591-bc8368fa0e60"`
}

type InventorySwagger5 struct {
	Sum float64 `json:"sum"`
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

	var price Inventory

	if err := components.NewDec(r, &price); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.SetPriceDB(price); err != nil {
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
	var inventory Inventory

	if err := components.NewDec(r, &inventory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.UpdateQuantity(inventory); err != nil {
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

	var discount NewInventoryDiscount

	if err := components.NewDec(r, &discount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.CreatingADiscount(discount); err != nil {
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
	var product WarehousePagination

	if err := components.NewDec(r, &product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	res, err := s.ListProductsByWarehouse(product)
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
	var product Inventory

	if err := components.NewDec(r, &product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	res, err := s.ListProduct(product)
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
	var count NewInventory

	if err := components.NewDec(r, &count); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	res, err := s.ListCount(count)
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
	var purchase NewInventory

	if err := components.NewDec(r, &purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.Purchase(purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
