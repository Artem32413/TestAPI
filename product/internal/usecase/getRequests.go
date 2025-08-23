package usecase

import (
	"product/pkg/errors"
	"product/pkg/headers"
	"product/pkg/requests"

	"context"
	"net/http"
	"time"
)

// DisplayAllProducts возвращает список всех товаров
// @Summary Получить все товары
// @Description Возвращает список всех товаров в системе без привязки к складам
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} structs.Products "Список товаров"
// @Failure 400 "Ошибка в запросе"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /products/all/ [get]
func (p *ProductHandler) DisplayAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	prod, err := p.svc.DisplayAllProductsLogic(p.logger, ctx)
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
