package appWarehouse

import (
	"apiGo/internal/warehouses/service"

	"go.uber.org/zap"
)

type WarehousesHandler struct {
	svc    *service.WarehousesService
	logger *zap.Logger
}

func New(svc *service.WarehousesService, logger *zap.Logger) *WarehousesHandler {
	return &WarehousesHandler{
		svc:    svc,
		logger: logger,
	}
}
