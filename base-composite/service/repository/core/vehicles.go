package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/gateway/controller"
	"github.com/iikmaulana/vehicle/base-composite/models"
	"github.com/iikmaulana/vehicle/base-composite/service"

	"github.com/iikmaulana/gateway/libs/helper/logger"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/packets"
	"time"
)

type vehiclesRepository struct {
	client packets.VehiclesClient
}

func NewVehiclesRepository(reg *controller.Registry) (res service.VehiclesRepo, serr serror.SError) {
	obj := vehiclesRepository{}
	i := 3

	for {
		if i <= 0 {
			serr = serror.Newc("Cannot connect to core Vehicles",
				"[repository][core][vehicles] while dialing core Vehicles")
			break
		}

		conn, err := reg.GetConnection("base-vehicle")
		if err != nil {
			logger.Warn("[repository][core][vehicles] while dial core Vehicles")
			time.Sleep(1 * time.Second)

			i--
			continue
		}
		obj.client = packets.NewVehiclesClient(conn)
		break
	}

	return obj, serr
}

func (v vehiclesRepository) AddVehiclesRepo(form models.VehVehiclesRequest) (vehicleId string, serr serror.SError) {

	byte, err := json.Marshal(form)
	if err != nil {
		return vehicleId, serror.NewFromError(err)
	}

	data := any.Any{
		Value: byte,
	}

	output, err := v.client.CreateVehicle(context.Background(), &packets.VehiclesRequest{
		Data: &data,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][vehicles][vehiclename:%d] while grpc AddVehiclesRepo: %+v", form.Name, err)
		logger.Err(serrFmt)
		return vehicleId, serror.NewFromErrorc(err, serrFmt)
	}
	if output.GetStatus() == 1 {
		vehicleId = string(output.GetData().Value)
	}

	return vehicleId, nil
}

func (v vehiclesRepository) UpdateVehiclesRepo(id string, form models.VehVehiclesUpdate) (serr serror.SError) {
	byte, err := json.Marshal(form)
	if err != nil {
		return serror.NewFromError(err)
	}

	data := any.Any{
		Value: byte,
	}

	output, err := v.client.UpdateVehicle(context.Background(), &packets.UpdateVehiclesRequest{
		VehiclesID: id,
		Data:       &data,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][vehicles][vehiclename:%d] while grpc UpdateVehicle: %+v", form.Name, err)
		logger.Err(serrFmt)
		return serror.NewFromErrorc(err, serrFmt)
	}
	if output.GetStatus() == 1 {

	}

	return nil
}

func (v vehiclesRepository) DeleteVehiclesByIdRepo(id string) (serr serror.SError) {
	_, err := v.client.DeleteVehiclesById(context.Background(), &packets.VehiclesRequestByID{
		VehiclesID: id,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][vehicles][vehiclesId:%d] while grpc DeleteVehiclesByIdRepo: %+v", id, err)
		logger.Err(serrFmt)
		return serror.NewFromErrorc(err, serrFmt)
	}

	return nil
}

func (v vehiclesRepository) GetVehiclesByIdRepo(vehiclesId string) (result []models.VehVehiclesResult, serr serror.SError) {

	output, err := v.client.GetVehiclesById(context.Background(), &packets.VehiclesRequestByID{
		VehiclesID: vehiclesId,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][vehicles][vehiclesId:%d] while grpc GetVehiclesByIdRepo: %+v", vehiclesId, err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromError(err)
		}
	}

	return result, nil
}

func (v vehiclesRepository) GetAllVehiclesRepo() (result models.VehListVehiclesResult, serr serror.SError) {
	output, err := v.client.GetVehiclesList(context.Background(), &packets.VehiclesRequest{})
	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][vehicles] while grpc GetAllVehiclesRepo: %+v", err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromError(err)
		}
	}

	return result, nil
}
