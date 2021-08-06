package impl

import (
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/vehicle/base/models"
	"github.com/iikmaulana/vehicle/base/service"
	"github.com/iikmaulana/vehicle/base/service/repository/query"
	"github.com/jmoiron/sqlx"
)

type settingRepository struct {
	DB *sqlx.DB
}

func NewSettingRepository(db *sqlx.DB) service.SettingRepo {
	return settingRepository{DB: db}
}

func (s settingRepository) AddSettingRepo(form models.VehSettingRequest) (serr serror.SError) {
	_, err := s.DB.Exec(query.CreateSetting,
		form.VehicleID,
		form.SpeedLimit,
		form.EngineKill,
		form.NotifOverspeed,
		form.NotifVehiclestop,
		form.NotifVehicleidle,
		form.NotifElectricwarning,
		form.NotifBatterybackup,
		form.NotifStarterkillon,
		form.NotifStarterkilloff,
		form.NotifTowing,
		form.NotifJamming,
		form.NotifIdling,
		form.NotifStop,
		form.StatStarterKill,
		form.NotifSimcardChange,
		form.NotifSignalPoor,
		form.NotifAccuDrop,
		form.CreateAt,
	)

	if err != nil {
		return serror.New("Can't exec query database")
	}

	return nil
}

func (s settingRepository) UpdateSettingRepo(id string, form models.VehSettingUpdate) (serr serror.SError) {
	panic("implement me")
}

func (s settingRepository) GetSettingByIdRepo(id string) (result []models.VehSettingRequest, serr serror.SError) {
	panic("implement me")
}

func (s settingRepository) GetAllSettingRepo() (result []models.VehSettingRequest, serr serror.SError) {
	panic("implement me")
}

func (s settingRepository) DeleteSettingByIdRepo(id string) (serr serror.SError) {
	panic("implement me")
}
