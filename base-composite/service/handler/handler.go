package handler

import (
	gateway "github.com/iikmaulana/gateway/service"
	"github.com/iikmaulana/vehicle/base-composite/service"
)

type gatewayHandler struct {
	service        *gateway.Service
	vehicleUsecase service.VehicleUsecase
}

func NewGatewayHandler(svc *gateway.Service,
	vehicleUsecase service.VehicleUsecase,
) {
	h := gatewayHandler{
		service:        svc,
		vehicleUsecase: vehicleUsecase,
	}

	h.initRoute()
}
