package service

import (
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/vehicle/base-composite/models"
)

type VehiclesRepo interface {
	AddVehiclesRepo(form models.VehVehiclesRequest) (vehicleId string, serr serror.SError)
	UpdateVehiclesRepo(id string, form models.VehVehiclesUpdate) (serr serror.SError)
	DeleteVehiclesByIdRepo(id string) (serr serror.SError)
	GetVehiclesByIdRepo(vehiclesId string) (result []models.VehVehiclesResult, serr serror.SError)
	GetAllVehiclesRepo() (result models.VehListVehiclesResult, serr serror.SError)
}
