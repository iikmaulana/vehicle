package handler

func (ox gatewayHandler) initRoute() {
	ox.service.Docs("docs/api-docs.yaml")

	ox.service.Strict(
		ox.service.POST("/vehicle", ox.createVehicle),
		ox.service.PUT("/vehicle/{id}", ox.updateVehicle),
		ox.service.DELETE("/vehicle/{id}", ox.deleteVehicleById),
	)

	ox.service.GET("/vehicle/view/{id}", ox.getVehicleById)
	ox.service.GET("/vehicle/list", ox.getAllVehicle)

}
