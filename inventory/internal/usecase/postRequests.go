package usecase

import (
	"inventory/internal/model/structs"
	"inventory/pkg/errors"
	"inventory/pkg/headers"
	"inventory/pkg/requests"

	"context"
	"net/http"
	"time"
)

// SetPrice Создание связи товара и склада
// @Summary Установить цену товара
// @Description Устанавливает цену для конкретного товара на указанном складе
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body structs.InventorySwagger1 true "Данные установки цены"
// @Success 200 "Цена успешно установлена"
// @Failure 400 "Неверные входные данные"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /inventory/price/ [post]
func (i *InventoryHandler) SetPrice(w http.ResponseWriter, r *http.Request) {

	var price structs.Inventory

	if err := requests.NewDec(r, &price); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := i.svc.SetPriceLogic(i.logger, ctx, price); err != nil {
		errors.HandleError(i.logger, w, err, http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(i.logger, w, []byte("Цена успешно установлена"))
}
