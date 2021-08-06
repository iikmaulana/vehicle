package query

const (
	CreateSetting = `INSERT INTO cleva.cleva_vehicle.veh_setting (vehicle_id, speed_limit, engine_kill, notif_overspeed, notif_vehiclestop, notif_vehicleidle, notif_electricwarning, notif_batterybackup, notif_starterkillon, notif_starterkilloff, notif_towing, notif_jamming, notif_idling, notif_stop, stat_starter_kill, notif_simcard_change, notif_signal_poor, notif_accu_drop, create_at)
					 VALUES
						($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`
	DeleteSetting = `update cleva.cleva_vehicle.veh_vehicle set delete_at = now(), status = 0 where id = $1`

	GetAllSetting = `select coalesce(vv.id, '')             as id,
						   coalesce(vv.name, '')            as name,
						   coalesce(vv.number_plate, '')    as number_plate,
						   coalesce(vv.chassis_number, '')  as chassis_number,
						   coalesce(vv.engine_number, '')   as engine_number,
						   coalesce(vv.device_imei, '')     as device_imei,
						   coalesce(vv.organization_id, '') as organization_id,
						   coalesce(vv.status, '')          as status,
						   coalesce(vv.created_at, '')      as created_at,
						   coalesce(vv.update_at, '')       as update_at,
						   coalesce(vv.delete_at, '')       as delete_at
					  from cleva.cleva_vehicle.veh_vehicle vv`

	GetSettingById = `select coalesce(vv.id, '')             as id,
						   coalesce(vv.name, '')            as name,
						   coalesce(vv.number_plate, '')    as number_plate,
						   coalesce(vv.chassis_number, '')  as chassis_number,
						   coalesce(vv.engine_number, '')   as engine_number,
						   coalesce(vv.device_imei, '')     as device_imei,
						   coalesce(vv.organization_id, '') as organization_id,
						   coalesce(vv.status, '')          as status,
						   coalesce(vv.created_at, '')      as created_at,
						   coalesce(vv.update_at, '')       as update_at,
						   coalesce(vv.delete_at, '')       as delete_at
					  from cleva.cleva_vehicle.veh_vehicle vv
					  where vv.id = ?`
)
