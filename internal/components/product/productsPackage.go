package product

import (
	"apiGo/internal/components"
	"go.uber.org/zap"
	"net/http"
)

// Products структура данных товара
// @Description Основная информация о товаре в системе
type Products struct {
	Product_id  string              `json:"product_id" example:"PRD-1001"`                        // Уникальный идентификатор товара
	Name        string              `json:"name" example:"Смартфон"`                              // Наименование товара
	Description string              `json:"description" example:"Флагманский смартфон 2023 года"` // Описание товара
	KeyValue    []map[string]string `json:"keyvalue,omitempty"`                                   // Характеристики товара (ключ-значение)
	Weight      string              `json:"weight,omitempty" example:"0.2"`                    // Вес товара
	Barcode     string              `json:"barcode,omitempty" example:"123456789012"`             // Штрих-код товара
}

type InventoryService struct {
	*components.Settings
}

// DisplayAllProducts возвращает список всех товаров
// @Summary Получить все товары
// @Description Возвращает список всех товаров в системе без привязки к складам
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} Products "Список товаров"
// @Failure 400 "Ошибка в запросе"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /products/all/ [get]
func (s *InventoryService) DisplayAllProducts(w http.ResponseWriter, r *http.Request) {
	prod, err := s.DisplayProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выведении всех товаров", zap.Error(err))
		return
	}

	jsData, err := components.NewMarshall(prod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в преобразовании JSON (Товары)", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error("Ошибка в выводе данных (Товары)", zap.Error(err))
		return
	}
}

// AddingNewProducts добавляет новый товар
// @Summary Добавить новый товар
// @Description Создает новую запись товара в системе
// @Tags Products
// @Accept json
// @Produce json
// @Param product body Products true "Данные нового товара"
// @Success 200 "Товар успешно добавлен"
// @Failure 400 "Неверные входные данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /products/add/ [post]
func (s *InventoryService) AddingNewProducts(w http.ResponseWriter, r *http.Request) {
	var products Products

	defer r.Body.Close()

	if err := components.NewDec(r, &products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.AdditionProducts(products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateProduct обновляет данные товара
// @Summary Обновить товар
// @Description Обновляет описание и характеристики существующего товара
// @Tags Products
// @Accept json
// @Produce json
// @Param product body Products true "Обновленные данные товара"
// @Success 200 "Товар успешно обновлен"
// @Failure 400 "Неверные входные данные"
// @Failure 404 "Товар не найден"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /products/update/ [put]
func (s *InventoryService) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var products Products

	if err := components.NewDec(r, &products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	if err := s.UpdateProd(products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
