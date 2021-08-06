package controller

import (
	"github.com/google/uuid"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/libs/utils/uttime"
	"github.com/iikmaulana/vehicle/base/models"
	"github.com/iikmaulana/vehicle/base/service"
	"time"
)

type vehicleUsecase struct {
	vehicleRepo service.VehiclesRepo
	settingRepo service.SettingRepo
}

func NewVehicleUsecase(vehicleRepo service.VehiclesRepo, settingRepo service.SettingRepo) service.VehiclesUsecase {
	return vehicleUsecase{vehicleRepo: vehicleRepo, settingRepo: settingRepo}
}

func (v vehicleUsecase) AddVehiclesUsecase(form models.VehVehiclesRequest) (id string, serr serror.SError) {
	if !v.vehicleRepo.CheckName(form.OrganizationID, form.Name) {
		return id, serror.New("Duplicate name")
	}

	if form.NumberPlate != "" {
		if !v.vehicleRepo.CheckNumberPlate(form.OrganizationID, form.NumberPlate) {
			return id, serror.New("Duplicate Number Plate")
		}
	}

	if form.ChassisNumber != "" {
		if !v.vehicleRepo.CheckChassisNumber(form.OrganizationID, form.ChassisNumber) {
			return id, serror.New("Duplicate Chassis Number")
		}
	}

	if form.EngineNumber != "" {
		if !v.vehicleRepo.CheckEngineNumber(form.OrganizationID, form.EngineNumber) {
			return id, serror.New("Duplicate Engine Number")
		}
	}

	if form.DeviceImei != "" {
		if !v.vehicleRepo.CheckDeviceImei(form.DeviceImei) {
			return id, serror.New("Duplicate Device Imei")
		}
	}

	if form.OrganizationID != "" {
		/*if v.vehicleRepo.CheckOrganizationId(form.OrganizationID) {
			return serror.New("Duplicate Organization ID")
		}*/
	}

	if form.Status != 0 {
		form.Status = 1
	}

	vehicles := models.VehVehicles{
		ID:             uuid.New().String(),
		Name:           form.Name,
		NumberPlate:    form.NumberPlate,
		ChassisNumber:  form.ChassisNumber,
		EngineNumber:   form.EngineNumber,
		DeviceImei:     form.DeviceImei,
		OrganizationID: uuid.New().String(),
		Status:         form.Status,
		CreatedAt:      uttime.ToString(uttime.DefaultDateTimeFormat, time.Now()),
	}

	id, serr = v.vehicleRepo.AddVehiclesRepo(vehicles)
	if serr != nil {
		return id, serr
	}

	formSetting := models.VehSettingRequest{
		VehicleID: id,
		CreateAt:  uttime.ToString(uttime.DefaultDateTimeFormat, time.Now()),
	}

	serr = v.settingRepo.AddSettingRepo(formSetting)
	if serr != nil {
		return id, serr
	}

	return id, nil
}

func (v vehicleUsecase) UpdateVehiclesUsecase(id string, form models.VehVehiclesUpdate) (serr serror.SError) {
	if !v.vehicleRepo.CheckName(form.OrganizationID, form.Name) {
		return serror.New("Duplicate name")
	}

	if form.NumberPlate != "" {
		if !v.vehicleRepo.CheckNumberPlate(form.OrganizationID, form.NumberPlate) {
			return serror.New("Duplicate Number Plate")
		}
	}

	if form.ChassisNumber != "" {
		if !v.vehicleRepo.CheckChassisNumber(form.OrganizationID, form.ChassisNumber) {
			return serror.New("Duplicate Chassis Number")
		}
	}

	if form.EngineNumber != "" {
		if !v.vehicleRepo.CheckEngineNumber(form.OrganizationID, form.EngineNumber) {
			return serror.New("Duplicate Engine Number")
		}
	}

	if form.DeviceImei != "" {
		if !v.vehicleRepo.CheckDeviceImei(form.DeviceImei) {
			return serror.New("Duplicate Device Imei")
		}
	}

	if form.OrganizationID != "" {
		/*if v.vehicleRepo.CheckOrganizationId(form.OrganizationID) {
			return serror.New("Duplicate Organization ID")
		}*/
	}

	if form.Status != 0 {
		form.Status = 1
	}

	vehicles := models.VehVehicles{
		Name:           form.Name,
		NumberPlate:    form.NumberPlate,
		ChassisNumber:  form.ChassisNumber,
		EngineNumber:   form.EngineNumber,
		DeviceImei:     form.DeviceImei,
		OrganizationID: form.OrganizationID,
		Status:         form.Status,
		UpdateAt:       uttime.ToString(uttime.DefaultDateTimeFormat, time.Now()),
	}

	serr = v.vehicleRepo.UpdateVehiclesRepo(id, vehicles)
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

	return serr
}

func (v vehicleUsecase) GetVehiclesByIdUsecase(id string) (result []models.VehVehiclesResult, serr serror.SError) {
	result, serr = v.vehicleRepo.GetVehiclesByIdRepo(id)
	if serr != nil {
		return result, serr
	}

	return result, serr
}

func (v vehicleUsecase) GetAllVehiclesUsecase(ndata int64, page int) (result models.VehListVehiclesResult, serr serror.SError) {
	vehicles, metas, serr := v.vehicleRepo.GetAllVehiclesRepo(ndata, page)
	if serr != nil {
		return result, serr
	}

	result.Result = metas
	result.Data = vehicles

	return result, serr
}
