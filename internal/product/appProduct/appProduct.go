package appProduct

import (
	model "apiGo/internal/product/model/structs"
	"apiGo/internal/product/service"
	"apiGo/pkg/errors"
	"apiGo/pkg/headers"
	"apiGo/pkg/requests"

	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type ProductHandler struct {
	svc    *service.ProductService
	logger *zap.Logger
}

func New(svc *service.ProductService, logger *zap.Logger) *ProductHandler {
	return &ProductHandler{
		svc:    svc,
		logger: logger,
	}
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
func (p *ProductHandler) DisplayAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	prod, err := p.svc.DisplayAllProductsLogic(ctx)
	if err != nil {
		errors.HandleError(p.logger, w, err, http.StatusBadRequest)
		return
	}

	jsData, err := requests.NewMarshall(prod)
	if err != nil {
		errors.HandleError(p.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(p.logger, w, jsData)
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
func (p *ProductHandler) AddingNewProducts(w http.ResponseWriter, r *http.Request) {
	var products model.Products

	defer r.Body.Close()

	if err := requests.NewDec(r, &products); err != nil {
		errors.HandleError(p.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := p.svc.AddingNewProductsLogic(ctx, products); err != nil {
		errors.HandleError(p.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(p.logger, w, []byte("Успешное добавление товара"))
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
func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var products model.Products

	if err := requests.NewDec(r, &products); err != nil {
		errors.HandleError(p.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := p.svc.UpdateProductLogic(ctx, products); err != nil {
		errors.HandleError(p.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(p.logger, w, []byte("Успешное обновление товара"))
}
