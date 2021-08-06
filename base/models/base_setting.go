package models

type VehSettingRequest struct {
	VehicleID            string `json:"vehicle_id" valid:"required"`
	SpeedLimit           string `json:"speed_limit"`
	EngineKill           int64  `json:"engine_kill"`
	NotifOverspeed       int64  `json:"notif_overspeed"`
	NotifVehiclestop     int64  `json:"notif_vehiclestop"`
	NotifVehicleidle     int64  `json:"notif_vehicleidle"`
	NotifElectricwarning int64  `json:"notif_electricwarning"`
	NotifBatterybackup   int64  `json:"notif_batterybackup"`
	NotifStarterkillon   int64  `json:"notif_starterkillon"`
	NotifStarterkilloff  int64  `json:"notif_starterkilloff"`
	NotifTowing          int64  `json:"notif_towing"`
	NotifJamming         int64  `json:"notif_jamming"`
	NotifIdling          int64  `json:"notif_idling"`
	NotifStop            int64  `json:"notif_stop"`
	StatStarterKill      int64  `json:"stat_starter_kill"`
	NotifSimcardChange   int64  `json:"notif_simcard_change"`
	NotifSignalPoor      int64  `json:"notif_signal_poor"`
	NotifAccuDrop        int64  `json:"notif_accu_drop"`
	CreateAt             string `json:"create_at"`
	UpdateAt             string `json:"update_at"`
	DeleteAt             string `json:"delete_at"`
}

type VehSettingUpdate struct {
	VehicleID            string `json:"vehicle_id,omitempty" valid:"required"`
	SpeedLimit           string `json:"speed_limit,omitempty"`
	EngineKill           int64  `json:"engine_kill,omitempty"`
	NotifOverspeed       int64  `json:"notif_overspeed,omitempty"`
	NotifVehiclestop     int64  `json:"notif_vehiclestop,omitempty"`
	NotifVehicleidle     int64  `json:"notif_vehicleidle,omitempty"`
	NotifElectricwarning int64  `json:"notif_electricwarning,omitempty"`
	NotifBatterybackup   int64  `json:"notif_batterybackup,omitempty"`
	NotifStarterkillon   int64  `json:"notif_starterkillon,omitempty"`
	NotifStarterkilloff  int64  `json:"notif_starterkilloff,omitempty"`
	NotifTowing          int64  `json:"notif_towing,omitempty"`
	NotifJamming         int64  `json:"notif_jamming,omitempty"`
	NotifIdling          int64  `json:"notif_idling,omitempty"`
	NotifStop            int64  `json:"notif_stop,omitempty"`
	StatStarterKill      int64  `json:"stat_starter_kill,omitempty"`
	NotifSimcardChange   int64  `json:"notif_simcard_change,omitempty"`
	NotifSignalPoor      int64  `json:"notif_signal_poor,omitempty"`
	NotifAccuDrop        int64  `json:"notif_accu_drop,omitempty"`
	CreateAt             string `json:"create_at,omitempty"`
	UpdateAt             string `json:"update_at,omitempty"`
	DeleteAt             string `json:"delete_at,omitempty"`
}

type VehSettingResult struct {
	VehicleID            string `json:"vehicle_id"`
	SpeedLimit           string `json:"speed_limit"`
	EngineKill           int64  `json:"engine_kill"`
	NotifOverspeed       int64  `json:"notif_overspeed"`
	NotifVehiclestop     int64  `json:"notif_vehiclestop"`
	NotifVehicleidle     int64  `json:"notif_vehicleidle"`
	NotifElectricwarning int64  `json:"notif_electricwarning"`
	NotifBatterybackup   int64  `json:"notif_batterybackup"`
	NotifStarterkillon   int64  `json:"notif_starterkillon"`
	NotifStarterkilloff  int64  `json:"notif_starterkilloff"`
	NotifTowing          int64  `json:"notif_towing"`
	NotifJamming         int64  `json:"notif_jamming"`
	NotifIdling          int64  `json:"notif_idling"`
	NotifStop            int64  `json:"notif_stop"`
	StatStarterKill      int64  `json:"stat_starter_kill"`
	NotifSimcardChange   int64  `json:"notif_simcard_change"`
	NotifSignalPoor      int64  `json:"notif_signal_poor"`
	NotifAccuDrop        int64  `json:"notif_accu_drop"`
	CreateAt             string `json:"create_at"`
	UpdateAt             string `json:"update_at"`
	DeleteAt             string `json:"delete_at"`
}
