package components

import "net/http"

// Health проверяет работоспособность сервиса
// @Summary Проверка здоровья сервиса
// @Description Проверяет, что сервис работает и доступен
// @Tags Сервис
// @Accept json
// @Produce plain
// @Success 200 "Сервис работает"
// @Failure 400 "Неверный запрос"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /api/health/ [get]
func (s *Settings) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}