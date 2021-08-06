package models

type VehVehiclesToGroupRequest struct {
	ID        string `json:"id"`
	GroupID   string `json:"group_id"`
	VehicleID string `json:"vehicle_id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
	DeleteAt  string `json:"delete_at"`
}

type VehVehiclesToGroupUpdate struct {
	GroupID   string `json:"group_id,omitempty"`
	VehicleID string `json:"vehicle_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdateAt  string `json:"update_at,omitempty"`
	DeleteAt  string `json:"delete_at,omitempty"`
}

type VehVehiclesToGroupResult struct {
	ID        string `json:"id"`
	GroupID   string `json:"group_id"`
	VehicleID string `json:"vehicle_id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
	DeleteAt  string `json:"delete_at"`
}
