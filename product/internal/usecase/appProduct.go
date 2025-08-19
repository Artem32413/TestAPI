package usecase

import (
	"product/internal/service"

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
