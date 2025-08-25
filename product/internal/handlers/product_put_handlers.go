package handlers

import (
	"product/internal/model/structs"
	"product/pkg/errors"
	"product/pkg/headers"
	"product/pkg/requests"

	"context"
	"net/http"
	"time"
)

// UpdateProduct обновляет данные товара
// @Summary Обновить товар
// @Description Обновляет описание и характеристики существующего товара
// @Tags Products
// @Accept json
// @Produce json
// @Param product body structs.Products true "Обновленные данные товара"
// @Success 200 "Товар успешно обновлен"
// @Failure 400 "Неверные входные данные"
// @Failure 404 "Товар не найден"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /products/update/ [put]
func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var products structs.Products

	if err := requests.NewDec(r, &products); err != nil {
		errors.HandleError(p.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := p.svc.UpdateProductLogic(p.logger, ctx, products); err != nil {
		errors.HandleError(p.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(p.logger, w, []byte("Успешное обновление товара"))
}
