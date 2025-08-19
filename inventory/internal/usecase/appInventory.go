package usecase

import (
	"inventory/internal/service"

	"go.uber.org/zap"
)

type InventoryHandler struct {
	svc    *service.InventoryService
	logger *zap.Logger
}

func New(svc *service.InventoryService, logger *zap.Logger) *InventoryHandler {
	return &InventoryHandler{
		svc:    svc,
		logger: logger,
	}
}


