package config

import (
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/vehicle/base-composite/controller"
	"github.com/iikmaulana/vehicle/base-composite/service/handler"
	"github.com/iikmaulana/vehicle/base-composite/service/repository/core"
)

func (cfg Config) InitService() serror.SError {

	vehiclesRepo, serr := core.NewVehiclesRepository(cfg.Registry)
	if serr != nil {
		return serr
	}

	vehicleUsecase := controller.NewVehicleUsecase(vehiclesRepo)

	handler.NewGatewayHandler(cfg.Gateway, vehicleUsecase)

	return nil
}
