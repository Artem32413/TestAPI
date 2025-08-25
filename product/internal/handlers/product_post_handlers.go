package handlers

import (
	model "product/internal/model/structs"
	"product/pkg/errors"
	"product/pkg/headers"
	"product/pkg/requests"

	"context"
	"net/http"
	"time"
)

// AddingNewProducts добавляет новый товар
// @Summary Добавить новый товар
// @Description Создает новую запись товара в системе
// @Tags Products
// @Accept json
// @Produce json
// @Param product body structs.Products true "Данные нового товара"
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

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := p.svc.AddingNewProductsLogic(p.logger, ctx, products); err != nil {
		errors.HandleError(p.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(p.logger, w, []byte("Успешное добавление товара"))
}
