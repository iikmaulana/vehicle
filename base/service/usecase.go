package service

import (
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/vehicle/base/models"
)

type VehiclesUsecase interface {
	AddVehiclesUsecase(form models.VehVehiclesRequest) (id string, serr serror.SError)
	UpdateVehiclesUsecase(id string, form models.VehVehiclesUpdate) (serr serror.SError)
	DeleteVehicleByIdUsecase(id string) (serr serror.SError)
	GetVehiclesByIdUsecase(id string) (result []models.VehVehiclesResult, serr serror.SError)
	GetAllVehiclesUsecase(ndata int64, page int) (result models.VehListVehiclesResult, serr serror.SError)
}

type VehicleToGroupUsecase interface {
	AddVehicleToGroupUsecase(form models.VehVehiclesToGroupRequest) (serr serror.SError)
	UpdateVehicleToGroupUsecase(id string, form models.VehVehiclesToGroupUpdate) (serr serror.SError)
	GetVehicleToGroupByIdUsecase(id string) (result []models.VehVehiclesToGroupRequest, serr serror.SError)
	GetAllVehicleToGroupUsecase() (result []models.VehVehiclesToGroupRequest, serr serror.SError)
	DeleteVehicleToGroupByIdUsecase(id string) (serr serror.SError)
}

type SettingUsecase interface {
	AddSettingUsecase(form models.VehSettingRequest) (serr serror.SError)
	UpdateSettingUsecase(id string, form models.VehSettingUpdate) (serr serror.SError)
	GetSettingByIdUsecase(id string) (result []models.VehSettingRequest, serr serror.SError)
	GetAllSettingUsecase() (result []models.VehSettingRequest, serr serror.SError)
	DeleteSettingByIdUsecase(id string) (serr serror.SError)
}

type GroupUsecase interface {
	AddGroupUsecase(form models.VehGroupRequest) (serr serror.SError)
	UpdateGroupUsecase(id string, form models.VehGroupUpdate) (serr serror.SError)
	GetGroupByIdUsecase(id string) (result []models.VehGroupResult, serr serror.SError)
	GetAllGroupUsecase() (result []models.VehGroupResult, serr serror.SError)
	DeleteGroupByIdUsecase(id string) (serr serror.SError)
}

type UserToGroupUsecase interface {
	AddUserToGroupUsecase(form models.VehUserToGroupRequest) (serr serror.SError)
	UpdateUserToGroupUsecase(id string, form models.VehUserToGroupUpdate) (serr serror.SError)
	GetUserToGroupByIdUsecase(id string) (result []models.VehUserToGroupResult, serr serror.SError)
	GetAllUserToGroupUsecase() (result []models.VehUserToGroupResult, serr serror.SError)
	DeleteUserToGroupByIdUsecase(id string) (serr serror.SError)
}
