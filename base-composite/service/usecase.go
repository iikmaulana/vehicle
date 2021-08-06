package service

import (
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/vehicle/base-composite/models"
)

type VehicleUsecase interface {
	AddVehicleUsecase(organizationID string, form models.VehVehiclesRequest) (serr serror.SError)
	UpdateVehicleUsecase(id string, form models.VehVehiclesUpdate) (serr serror.SError)
	DeleteVehicleByIdUsecase(id string) (serr serror.SError)
	GetVehicleByIdUsecase(id string) (result []models.VehVehiclesResult, serr serror.SError)
	GetAllVehicleUsecase() (result models.VehListVehiclesResult, serr serror.SError)
}
