package controller

import (
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/vehicle/base-composite/models"
	"github.com/iikmaulana/vehicle/base-composite/service"
)

type vehicleUsecase struct {
	vehicleRepo service.VehiclesRepo
}

func NewVehicleUsecase(vehicleRepo service.VehiclesRepo) service.VehicleUsecase {
	return vehicleUsecase{vehicleRepo: vehicleRepo}
}

func (v vehicleUsecase) AddVehicleUsecase(organizationID string, form models.VehVehiclesRequest) (serr serror.SError) {

	_, serr = v.vehicleRepo.AddVehiclesRepo(form)
	if serr != nil {
		return serr
	}

	return nil
}

func (v vehicleUsecase) UpdateVehicleUsecase(id string, form models.VehVehiclesUpdate) (serr serror.SError) {
	serr = v.vehicleRepo.UpdateVehiclesRepo(id, form)
	if serr != nil {
		return serr
	}
	return nil
}

func (v vehicleUsecase) DeleteVehicleByIdUsecase(id string) (serr serror.SError) {
	serr = v.vehicleRepo.DeleteVehiclesByIdRepo(id)
	if serr != nil {
		return serr
	}
	return nil
}

func (v vehicleUsecase) GetVehicleByIdUsecase(id string) (result []models.VehVehiclesResult, serr serror.SError) {
	result, serr = v.vehicleRepo.GetVehiclesByIdRepo(id)
	if serr != nil {
		return result, serr
	}
	return result, serr
}

func (v vehicleUsecase) GetAllVehicleUsecase() (result models.VehListVehiclesResult, serr serror.SError) {
	result, serr = v.vehicleRepo.GetAllVehiclesRepo()
	if serr != nil {
		return result, serr
	}
	return result, serr
}
