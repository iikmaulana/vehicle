package grpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/packets"
	"github.com/iikmaulana/vehicle/base/models"
	"github.com/iikmaulana/vehicle/base/service"
)

type Handler struct {
	vehicleUsecase service.VehiclesUsecase
}

func NewVehicleHandler(vehicleUsecase service.VehiclesUsecase) Handler {
	return Handler{
		vehicleUsecase: vehicleUsecase}
}

func (handler Handler) CreateVehicle(ctx context.Context, form *packets.VehiclesRequest) (output *packets.OutputVehicles, err error) {
	output = &packets.OutputVehicles{
		Status: 0,
	}

	request := models.VehVehiclesRequest{}
	errx := json.Unmarshal(form.GetData().Value, &request)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	id, serr := handler.vehicleUsecase.AddVehiclesUsecase(request)
	if serr != nil {
		return output, serr
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: []byte(id),
	}

	return output, nil
}

func (handler Handler) UpdateVehicle(ctx context.Context, form *packets.UpdateVehiclesRequest) (output *packets.OutputVehicles, err error) {
	output = &packets.OutputVehicles{
		Status: 0,
	}

	request := models.VehVehiclesUpdate{}
	errx := json.Unmarshal(form.GetData().Value, &request)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	serr := handler.vehicleUsecase.UpdateVehiclesUsecase(form.VehiclesID, request)
	if serr != nil {
		return output, serr
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: []byte(""),
	}

	return output, nil
}

func (handler Handler) GetVehiclesById(ctx context.Context, form *packets.VehiclesRequestByID) (output *packets.OutputVehicles, err error) {
	output = &packets.OutputVehicles{
		Status: 0,
	}

	result, serr := handler.vehicleUsecase.GetVehiclesByIdUsecase(form.VehiclesID)
	if serr != nil {
		return output, serr
	}

	byte, err := json.Marshal(result)
	if err != nil {
		return output, serror.NewFromError(err)
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: byte,
	}

	return output, nil
}

func (handler Handler) GetVehiclesList(ctx context.Context, form *packets.VehiclesRequest) (output *packets.OutputVehicles, err error) {
	output = &packets.OutputVehicles{
		Status: 0,
	}

	result, serr := handler.vehicleUsecase.GetAllVehiclesUsecase(10, 1)
	if serr != nil {
		return output, serr
	}

	result.Result.Data = result.Data

	b, err := json.Marshal(&result.Result)
	if err != nil {
		return output, serror.NewFromError(err)
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: b,
	}

	return output, nil
}

func (handler Handler) DeleteVehiclesById(ctx context.Context, form *packets.VehiclesRequestByID) (output *packets.OutputVehicles, err error) {
	output = &packets.OutputVehicles{
		Status: 0,
	}

	serr := handler.vehicleUsecase.DeleteVehicleByIdUsecase(form.VehiclesID)
	if serr != nil {
		return output, serr
	}

	return output, nil
}
