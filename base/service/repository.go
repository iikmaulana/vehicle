package service

import (
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/vehicle/base/models"
)

type VehiclesRepo interface {
	AddVehiclesRepo(form models.VehVehicles) (id string, serr serror.SError)
	UpdateVehiclesRepo(id string, form models.VehVehicles) (serr serror.SError)
	DeleteVehiclesByIdRepo(id string) (serr serror.SError)
	GetVehiclesByIdRepo(vehiclesId string) (result []models.VehVehiclesResult, serr serror.SError)
	GetAllVehiclesRepo(ndata int64, page int) (result []models.VehVehiclesResult, metas models.FormMetaData, serr serror.SError)
	CheckName(organizationId string, name string) (result bool)
	CheckNumberPlate(organizationId string, numberPlate string) (result bool)
	CheckChassisNumber(organizationId string, chassisNumber string) (result bool)
	CheckEngineNumber(organizationId string, engineNumber string) (result bool)
	CheckDeviceImei(deviceImei string) (result bool)
	CheckOrganizationId(organizationId string) (result bool)
}

type UserToGroupRepo interface {
	AddUserToGroupRepo(form models.VehUserToGroupRequest) (serr serror.SError)
	UpdateUserToGroupRepo(id string, form models.VehUserToGroupUpdate) (serr serror.SError)
	GetUserToGroupByIdRepo(id string) (result []models.VehUserToGroupResult, serr serror.SError)
	GetAllUserToGroupRepo() (result []models.VehUserToGroupResult, serr serror.SError)
	DeleteUserToGroupByIdRepo(id string) (serr serror.SError)
}

type GroupRepo interface {
	AddGroupRepo(form models.VehGroupRequest) (serr serror.SError)
	UpdateGroupRepo(id string, form models.VehGroupUpdate) (serr serror.SError)
	GetGroupByIdRepo(id string) (result []models.VehGroupResult, serr serror.SError)
	GetAllGroupRepo() (result []models.VehGroupResult, serr serror.SError)
	DeleteGroupByIdRepo(id string) (serr serror.SError)
}

type VehicleToGroupRepo interface {
	AddVehicleToGroupRepo(form models.VehVehiclesToGroupRequest) (serr serror.SError)
	UpdateVehicleToGroupRepo(id string, form models.VehVehiclesToGroupUpdate) (serr serror.SError)
	GetVehicleToGroupByIdRepo(id string) (result []models.VehVehiclesToGroupRequest, serr serror.SError)
	GetAllVehicleToGroupRepo() (result []models.VehVehiclesToGroupRequest, serr serror.SError)
	DeleteVehicleToGroupByIdRepo(id string) (serr serror.SError)
}

type SettingRepo interface {
	AddSettingRepo(form models.VehSettingRequest) (serr serror.SError)
	UpdateSettingRepo(id string, form models.VehSettingUpdate) (serr serror.SError)
	GetSettingByIdRepo(id string) (result []models.VehSettingRequest, serr serror.SError)
	GetAllSettingRepo() (result []models.VehSettingRequest, serr serror.SError)
	DeleteSettingByIdRepo(id string) (serr serror.SError)
}
