package handlers

import (
	"inventory/internal/service"

	"go.uber.org/zap"
)

type InventoryHandler struct {
	svc    *service.InventoryService
	logger *zap.Logger
}

func NewInventoryHandlers(svc *service.InventoryService, logger *zap.Logger) *InventoryHandler {
	return &InventoryHandler{
		svc:    svc,
		logger: logger,
	}
}


