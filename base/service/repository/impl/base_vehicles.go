package impl

import (
	"fmt"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/vehicle/base/models"
	"github.com/iikmaulana/vehicle/base/service"
	"github.com/iikmaulana/vehicle/base/service/helper"
	"github.com/iikmaulana/vehicle/base/service/repository/query"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
	"strings"
)

type vehiclesRepository struct {
	DB *sqlx.DB
}

func NewVehiclesRepository(db *sqlx.DB) service.VehiclesRepo {
	return vehiclesRepository{DB: db}
}

func (v vehiclesRepository) AddVehiclesRepo(form models.VehVehicles) (id string, serr serror.SError) {

	err := v.DB.QueryRow(query.CreateVehicle,
		form.ID,
		form.Name,
		form.NumberPlate,
		form.ChassisNumber,
		form.EngineNumber,
		form.DeviceImei,
		form.OrganizationID,
		form.Status,
		form.CreatedAt,
	).Scan(&id)

	if err != nil {
		return id, serror.New("Can't exec query database")
	}

	return id, nil
}

func (v vehiclesRepository) UpdateVehiclesRepo(id string, form models.VehVehicles) (serr serror.SError) {

	var dynamicUpdate []string
	n := 0
	x := reflect.ValueOf(form)
	num := x.NumField()
	if num <= 0 {
		return serr
	}
	for i := 0; i < num; i++ {
		coloumn := x.Type().Field(i).Tag.Get("db")
		t := x.Type().Field(i).Type
		exist := x.Field(i).Interface() != reflect.Zero(t).Interface()
		if exist {
			v := fmt.Sprint(x.Field(i).Interface())
			n = n + 1
			q := coloumn + ` = ` + `'` + v + `'`
			dynamicUpdate = append(dynamicUpdate, q)
		}
	}

	prefix := ` WHERE id = $1`
	queryVehicleID := prefix
	prefix = `UPDATE cleva.cleva_vehicle.veh_vehicle SET `
	queryUpdateVehicle := prefix + strings.Join(dynamicUpdate, ",") + queryVehicleID

	_, err := v.DB.Exec(queryUpdateVehicle, id)

	if err != nil {
		log.Printf("failed %+v", err)
		return serror.New("Can't exec query database")
	}

	return nil
}

func (v vehiclesRepository) DeleteVehiclesByIdRepo(id string) (serr serror.SError) {
	_, err := v.DB.Exec(query.DeleteVehicle, id)
	if err != nil {
		return serror.New("Can't exec query database")
	}

	return nil
}

func (v vehiclesRepository) GetVehiclesByIdRepo(vehiclesId string) (result []models.VehVehiclesResult, serr serror.SError) {
	rows, err := v.DB.Queryx(query.GetVehicleById, vehiclesId)
	if err != nil {
		return result, serror.New("Can't exec query database")
	}

	defer rows.Close()
	for rows.Next() {
		tmp := models.VehVehiclesResult{}
		if err := rows.StructScan(&tmp); err != nil {
			return nil, serror.New("Failed scan for struct model")
		}
		result = append(result, tmp)
	}

	return result, nil
}

func (v vehiclesRepository) GetAllVehiclesRepo(ndata int64, page int) (result []models.VehVehiclesResult, metas models.FormMetaData, serr serror.SError) {

	rows, err := v.DB.Queryx(query.GetAllVehicle)
	if err != nil {
		return result, metas, serror.New("Can't exec query database")
	}

	defer rows.Close()
	for rows.Next() {
		temp := models.VehVehiclesResult{}
		if err := rows.StructScan(&temp); err != nil {
			return result, metas, serror.New("Failed scan for struct model")
		}
		result = append(result, temp)
	}

	var condition string

	var withPagination = true
	if page == 0 {
		withPagination = false
	}

	var offset int
	var paginate models.FormMetaData

	if withPagination == true {
		paginate, offset = helper.Paginate(ndata, page, 1)
		condition = condition + fmt.Sprintf(" group by lrr.loan_request_restructure_id, l.payment_type, lrsv.loan_id order by lrr.loan_request_restructure_id desc limit 1 offset %v ", offset)
	} else {
		condition = condition + " group by lrr.loan_request_restructure_id, l.payment_type, lrsv.loan_id order by lrr.loan_request_restructure_id desc "
	}

	result = result
	metas = paginate
	return result, metas, nil
}

func (v vehiclesRepository) CheckName(organizationId string, name string) (result bool) {
	if err := v.DB.QueryRow(query.QueryCheckName, organizationId, name).Scan(&result); err != nil {
		return !result
	}
	return !result
}

func (v vehiclesRepository) CheckNumberPlate(organizationId string, numberPlate string) (result bool) {
	if err := v.DB.QueryRow(query.QueryCheckNumberPlate, organizationId, numberPlate).Scan(&result); err != nil {
		return !result
	}
	return !result
}

func (v vehiclesRepository) CheckChassisNumber(organizationId string, chassisNumber string) (result bool) {
	if err := v.DB.QueryRow(query.QueryCheckChassisNumber, organizationId, chassisNumber).Scan(&result); err != nil {
		return !result
	}
	return !result
}

func (v vehiclesRepository) CheckEngineNumber(organizationId string, engineNumber string) (result bool) {
	if err := v.DB.QueryRow(query.QueryCheckEngineNumber, organizationId, engineNumber).Scan(&result); err != nil {
		return !result
	}
	return !result
}

func (v vehiclesRepository) CheckDeviceImei(deviceImei string) (result bool) {
	if err := v.DB.QueryRow(query.QueryCheckDeviceImei, deviceImei).Scan(&result); err != nil {
		return !result
	}
	return !result
}

func (v vehiclesRepository) CheckOrganizationId(organizationId string) (result bool) {
	if err := v.DB.QueryRow(query.QueryCheckOrganizationId, organizationId).Scan(&result); err != nil {
		return !result
	}
	return !result
}
