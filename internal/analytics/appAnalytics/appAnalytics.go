package appAnalytics

import (
	"apiGo/internal/analytics/model/structs"
	"apiGo/internal/analytics/service"
	"apiGo/pkg/errors"
	"apiGo/pkg/headers"
	"apiGo/pkg/requests"

	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type AnalyticsHandler struct {
	svc    *service.AnalyticsService
	logger *zap.Logger
}

func New(svc *service.AnalyticsService, logger *zap.Logger) *AnalyticsHandler {
	return &AnalyticsHandler{
		svc:    svc,
		logger: logger,
	}
}

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
	var a model.Analytics

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

	headers.HeaderWithSub(s.logger, w, jsData)
}

// Top обрабатывает запрос топовых складов
// @Summary Получить топ складов по выручке
// @Description Возвращает список складов с наибольшей выручкой
// @Tags Analytics
// @Accept json
// @Produce json
// @Success 200 {array} TopAnalytics "Список топовых складов"
// @Failure 400 "Ошибка в запросе"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /analytics/top/ [get]
func (s *AnalyticsHandler) AnalyticsTop(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	res, err := s.svc.AnalyticsTopLogic(ctx)
	if err != nil {
		errors.HandleError(s.logger, w, err, http.StatusBadRequest)
		return
	}

	jsData, err := requests.NewMarshall(res)
	if err != nil {
		errors.HandleError(s.logger, w, fmt.Errorf("Ошибка в преобразовании JSON (Склады)"), http.StatusBadRequest)
		return
	}

	headers.HeaderWithSub(s.logger, w, jsData)
}
