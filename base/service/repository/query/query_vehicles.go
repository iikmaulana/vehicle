package query

const (
	CreateVehicle = `INSERT INTO cleva.cleva_vehicle.veh_vehicle
						(id, name, number_plate, chassis_number, engine_number, device_imei, organization_id, status, created_at)
					 VALUES
						($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`
	DeleteVehicle = `update cleva.cleva_vehicle.veh_vehicle set delete_at = now(), status = 0 where id = $1`

	GetAllVehicle = `select coalesce(vv.id::text, '')              as id,
						   coalesce(vv.name, '')                  as name,
						   coalesce(vv.number_plate, '')          as number_plate,
						   coalesce(vv.chassis_number, '')        as chassis_number,
						   coalesce(vv.engine_number, '')         as engine_number,
						   coalesce(vv.device_imei, '')           as device_imei,
						   coalesce(vv.organization_id::text, '') as organization_id,
						   coalesce(vv.status, 0)                 as status,
						   coalesce(vv.created_at::text, '')      as created_at,
						   coalesce(vv.update_at::text, '')       as update_at,
						   coalesce(vv.delete_at::text, '')       as delete_at
					from cleva.cleva_vehicle.veh_vehicle vv
					where vv.delete_at is null`

	GetVehicleById = `select coalesce(vv.id::text, '')              as id,
						   coalesce(vv.name, '')                  as name,
						   coalesce(vv.number_plate, '')          as number_plate,
						   coalesce(vv.chassis_number, '')        as chassis_number,
						   coalesce(vv.engine_number, '')         as engine_number,
						   coalesce(vv.device_imei, '')           as device_imei,
						   coalesce(vv.organization_id::text, '') as organization_id,
						   coalesce(vv.status, 0)                 as status,
						   coalesce(vv.created_at::text, '')      as created_at,
						   coalesce(vv.update_at::text, '')       as update_at,
						   coalesce(vv.delete_at::text, '')       as delete_at
					from cleva.cleva_vehicle.veh_vehicle vv
					where vv.name = $1 or vv.number_plate = $1 or vv.chassis_number = $1 and vv.delete_at is null`

	QueryCheckName           = `select exists ( select id from cleva.cleva_vehicle.veh_vehicle where organization_id = $1 and name = $2 and delete_at is null limit 1) as value`
	QueryCheckNumberPlate    = `select exists ( select id from cleva.cleva_vehicle.veh_vehicle where organization_id = $1 and number_plate = $2 and delete_at is null limit 1) as value`
	QueryCheckChassisNumber  = `select exists ( select id from cleva.cleva_vehicle.veh_vehicle where organization_id = $1 and chassis_number = $2 and delete_at is null limit 1) as value`
	QueryCheckEngineNumber   = `select exists ( select id from cleva.cleva_vehicle.veh_vehicle where organization_id = $1 and engine_number = $2 and delete_at is null limit 1) as value`
	QueryCheckDeviceImei     = `select exists ( select id from cleva.cleva_vehicle.veh_vehicle where device_imei = $1 and delete_at is null limit 1) as value`
	QueryCheckOrganizationId = `select exists ( select id from cleva.cleva_vehicle.veh_vehicle where organization_id = $1 and delete_at is null limit 1) as value`
)
