package usecase

import (
	"analytics/internal/model/structs"
	"analytics/pkg/errors"
	"analytics/pkg/headers"
	"analytics/pkg/requests"

	"context"
	"fmt"
	"net/http"
	"time"
)

// AnalyticsAll обрабатывает запрос аналитики продаж
// @Summary Получить аналитику продаж
// @Description Возвращает аналитику продаж по указанным параметрам
// @Tags Analytics
// @Accept json
// @Produce json
// @Param request body AnalyticsSwagger true "Параметры запроса аналитики"
// @Success 200 {array} Analytics "Успешный ответ"
// @Failure 400 "Ошибка в запросе"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /analytics/ [post]
func (s *AnalyticsHandler) AnalyticsAll(w http.ResponseWriter, r *http.Request) {
	var a structs.Analytics

	if err := requests.NewDec(r, &a); err != nil {
		errors.HandleError(s.logger, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	res, err := s.svc.AnalyticsAllLogic(ctx, a)
	if err != nil {
		errors.HandleError(s.logger, w, err, http.StatusBadRequest)
		return
	}

	jsData, err := requests.NewMarshall(res)
	if err != nil {
		errors.HandleError(s.logger, w, fmt.Errorf("Ошибка в преобразовании JSON (Склады)"), http.StatusBadRequest)
		return
	}

	headers.HeaderWithText(s.logger, w, jsData)
}
