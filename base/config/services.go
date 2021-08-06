package config

import (
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/packets"
	"github.com/iikmaulana/vehicle/base/controller"
	"github.com/iikmaulana/vehicle/base/service/grpc"
	"github.com/iikmaulana/vehicle/base/service/repository/impl"
)

func (cfg Config) InitService() serror.SError {

	vehicleRepo := impl.NewVehiclesRepository(cfg.DB)
	settingRepo := impl.NewSettingRepository(cfg.DB)

	vehicleUsecase := controller.NewVehicleUsecase(vehicleRepo, settingRepo)

	packets.RegisterVehiclesServer(cfg.Server.Instance(), grpc.NewVehicleHandler(vehicleUsecase))

	return nil
}
