package usecase

import (
	"analytics/pkg/errors"
	"analytics/pkg/headers"
	"analytics/pkg/requests"

	"context"
	"fmt"
	"net/http"
	"time"
)

// Top обрабатывает запрос топовых складов
// @Summary Получить топ складов по выручке
// @Description Возвращает список складов с наибольшей выручкой
// @Tags Analytics
// @Accept json
// @Produce json
// @Success 200 {array} structs.TopAnalytics "Список топовых складов"
// @Failure 400 "Ошибка в запросе"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /analytics/top/ [get]
func (s *AnalyticsHandler) AnalyticsTop(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	res, err := s.svc.AnalyticsTopLogic(s.logger, ctx)
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
