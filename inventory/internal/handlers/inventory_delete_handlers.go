package handlers

import (
	"inventory/internal/model/structs"
	"inventory/pkg/errors"
	"inventory/pkg/headers"
	"inventory/pkg/requests"

	"context"
	"net/http"
	"time"
)

// PurchaseProduct обрабатывает покупку товаров
// @Summary Покупка товаров
// @Description Обрабатывает покупку товаров и уменьшает их количество на складе
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body structs.NewInventory true "Данные о покупке"
// @Success 200 "Покупка успешно обработана"
// @Failure 400 "Неверные входные данные или недостаточно товара"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/purchase/ [delete]
func (i *InventoryHandler) PurchaseProduct(w http.ResponseWriter, r *http.Request) {
	var purchase structs.NewInventory

	if err := requests.NewDec(r, &purchase); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := i.svc.PurchaseProductLogic(i.logger, ctx, purchase); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(i.logger, w, []byte("Покупка успешно обработана"))
}
