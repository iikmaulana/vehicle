package handler

import (
	"encoding/json"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/vehicle/base-composite/models"
	"net/http"

	response "github.com/iikmaulana/gateway/models"
	gateway "github.com/iikmaulana/gateway/service"
)

func (ox gatewayHandler) createVehicle(ctx *gateway.Context) gateway.Result {

	level := ctx.AuthorizationInfo().IsOrgAdmin
	organizationId := ctx.AuthorizationInfo().OrganizationId

	if level != 1 && level != 2 && level != 3 {
		serr := serror.New("Token access denied")
		return ctx.JSONResponse(http.StatusUnauthorized, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}
	form := models.VehVehiclesRequest{}
	err := json.Unmarshal(ctx.RawBody(), &form)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusBadRequest, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	err = ox.vehicleUsecase.AddVehicleUsecase(organizationId, form)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusInternalServerError, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	return ctx.JSONResponse(http.StatusOK, response.ResponseBody{
		Appid:      "",
		Svcid:      ox.service.Key(),
		Controller: ctx.Path(),
		Action:     ctx.Method(),
	})
}

func (ox gatewayHandler) updateVehicle(ctx *gateway.Context) gateway.Result {
	level := ctx.AuthorizationInfo().IsOrgAdmin

	if level != 1 && level != 2 && level != 3 {
		serr := serror.New("Token access denied")
		return ctx.JSONResponse(http.StatusUnauthorized, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}
	id := ctx.Parameter("id")
	var form models.VehVehiclesUpdate
	err := json.Unmarshal(ctx.RawBody(), &form)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusBadRequest, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	err = ox.vehicleUsecase.UpdateVehicleUsecase(id, form)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusInternalServerError, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	return ctx.JSONResponse(http.StatusOK, response.ResponseBody{
		Appid:      "",
		Svcid:      ox.service.Key(),
		Controller: ctx.Path(),
		Action:     ctx.Method(),
	})
}

func (ox gatewayHandler) getVehicleById(ctx *gateway.Context) gateway.Result {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	id := ctx.Parameter("id")
	result, err := ox.vehicleUsecase.GetVehicleByIdUsecase(id)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusInternalServerError, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	return ctx.JSONResponse(http.StatusOK, response.ResponseBody{
		Appid:      "",
		Svcid:      ox.service.Key(),
		Controller: ctx.Path(),
		Action:     ctx.Method(),
		Result:     result,
	})
}

func (ox gatewayHandler) getAllVehicle(ctx *gateway.Context) gateway.Result {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	result, err := ox.vehicleUsecase.GetAllVehicleUsecase()
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusInternalServerError, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	return ctx.JSONResponse(http.StatusOK, response.ResponseBody{
		Appid:      "",
		Svcid:      ox.service.Key(),
		Controller: ctx.Path(),
		Action:     ctx.Method(),
		Result:     result,
	})
}

func (ox gatewayHandler) deleteVehicleById(ctx *gateway.Context) gateway.Result {
	level := ctx.AuthorizationInfo().IsOrgAdmin

	if level != 1 && level != 2 && level != 3 {
		serr := serror.New("Token access denied")
		return ctx.JSONResponse(http.StatusUnauthorized, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}
	id := ctx.Parameter("id")
	err := ox.vehicleUsecase.DeleteVehicleByIdUsecase(id)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusInternalServerError, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	return ctx.JSONResponse(http.StatusOK, response.ResponseBody{
		Appid:      "",
		Svcid:      ox.service.Key(),
		Controller: ctx.Path(),
		Action:     ctx.Method(),
	})
}
