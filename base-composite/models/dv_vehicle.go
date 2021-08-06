package models

type VehVehicles struct {
	ID             string `json:"id"`
	Name           string `db:"name" json:"name"`
	NumberPlate    string `db:"number_plate" json:"number_plate"`
	ChassisNumber  string `db:"chassis_number" json:"chassis_number"`
	EngineNumber   string `db:"engine_number" json:"engine_number"`
	DeviceImei     string `db:"device_imei" json:"device_imei"`
	OrganizationID string `db:"organization_id" json:"organization_id"`
	Status         int64  `db:"status" json:"status"`
	CreatedAt      string `db:"created_at" json:"created_at"`
	UpdateAt       string `db:"update_at" json:"update_at"`
	DeleteAt       string `db:"delete_at" json:"delete_at"`
}

type VehVehiclesRequest struct {
	Name           string `json:"name" valid:"required"`
	NumberPlate    string `json:"number_plate"`
	ChassisNumber  string `json:"chassis_number"`
	EngineNumber   string `json:"engine_number"`
	DeviceImei     string `json:"device_imei"`
	OrganizationID string `json:"organization_id" valid:"required"`
	Status         int64  `json:"status"`
}

type VehVehiclesUpdate struct {
	Name           string `json:"name,omitempty" valid:"required"`
	NumberPlate    string `json:"number_plate,omitempty"`
	ChassisNumber  string `json:"chassis_number,omitempty"`
	EngineNumber   string `json:"engine_number,omitempty"`
	DeviceImei     string `json:"device_imei,omitempty"`
	OrganizationID string `json:"organization_id,omitempty"`
	Status         int64  `json:"status,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdateAt       string `json:"update_at,omitempty"`
	DeleteAt       string `json:"delete_at,omitempty"`
}

type VehListVehiclesResult struct {
	NData     int                 `json:"NData"`
	Page      int64               `json:"Page"`
	TotalPage int64               `json:"TotalPage"`
	TotalData int                 `json:"TotalData"`
	From      int                 `json:"From"`
	To        int                 `json:"To"`
	Data      []VehVehiclesResult `json:"Data"`
}

type VehVehiclesResult struct {
	ID             string `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	NumberPlate    string `db:"number_plate" json:"number_plate"`
	ChassisNumber  string `db:"chassis_number" json:"chassis_number"`
	EngineNumber   string `db:"engine_number" json:"engine_number"`
	DeviceImei     string `db:"device_imei" json:"device_imei"`
	OrganizationID string `db:"organization_id" json:"organization_id"`
	Status         int64  `db:"status" json:"status"`
	CreatedAt      string `db:"created_at" json:"created_at"`
	UpdateAt       string `db:"update_at" json:"update_at"`
	DeleteAt       string `db:"delete_at" json:"delete_at"`
}
