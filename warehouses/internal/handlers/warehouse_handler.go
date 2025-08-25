package handlers

import (
	"warehouses/internal/service"

	"go.uber.org/zap"
)

type WarehousesHandler struct {
	svc    *service.WarehousesService
	logger *zap.Logger
}

func NewWarehouseHandler(svc *service.WarehousesService, logger *zap.Logger) *WarehousesHandler {
	return &WarehousesHandler{
		svc:    svc,
		logger: logger,
	}
}
