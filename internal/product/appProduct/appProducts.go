package appProducts

import (
	"apiGo/internal/components"
	"apiGo/internal/product/config/settings"
	"apiGo/internal/product/model/interfaces"
	model "apiGo/internal/product/model/structs"
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)


type InventoryService struct {
	*settings.Settings
}

var st struct{
	interfaces.ProductsRepo
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	prod, err := st.DisplayProducts(ctx)
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
	var products model.Products

	defer r.Body.Close()

	if err := components.NewDec(r, &products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := st.AdditionProducts(ctx, products); err != nil {
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
	var products model.Products

	if err := components.NewDec(r, &products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := st.UpdateProd(ctx, products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
